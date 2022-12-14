package mch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MchToken struct {
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
}

type MchSession struct {
	Config *MchConfig
	Token  *MchToken
}

func Login(clientId string, clientSecret string, username string, password string) (*MchSession, error) {
	config, err := GetConfiguration()
	if err != nil {
		return nil, err
	}

	authUrl := config.GetString("cloud.service.urls", "service.auth0.url")
	authUrl = fmt.Sprintf("%s/oauth/token", authUrl)

	req := map[string]string{
		"grant_type":    "http://auth0.com/oauth/grant-type/password-realm",
		"realm":         "Username-Password-Authentication",
		"audience":      "mycloud.com",
		"username":      username,
		"password":      password,
		"scope":         "openid offline_access nas_read_write nas_read_only user_read device_read",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &MchSession{Config: config}, err
	}

	httpClient := http.Client{}
	res, err := httpClient.Post(
		authUrl,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return &MchSession{Config: config}, err
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return &MchSession{Config: config}, fmt.Errorf("status code %d has been received from %s", res.StatusCode, res.Request.URL)
	}

	var token MchToken
	err = json.NewDecoder(res.Body).Decode(&token)

	if err != nil {
		return &MchSession{Config: config}, err
	}

	return &MchSession{Config: config, Token: &token}, nil
}

func Relogin(clientId string, clientSecret string, session *MchSession) error {

	req := map[string]string{
		"audience":      "mycloud.com",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": session.Token.RefreshToken,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	res, err := httpClient.Post(
		fmt.Sprintf("%s/oauth/token", session.Config.GetString("cloud.service.urls", "service.auth0.url")),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return fmt.Errorf("status code %d has been received from %s", res.StatusCode, res.Request.URL)
	}

	b, _ := io.ReadAll(res.Body)
	log.Print(string(b))

	return nil
}

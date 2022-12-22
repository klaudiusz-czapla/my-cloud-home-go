package mch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/serde"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
)

type MchProxy struct {
	HttpClient *http.Client
	Session    *MchSession
}

func Login(clientId string, clientSecret string, username string, password string) (*MchProxy, error) {
	config, err := serde.GetConfiguration()
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
		return &MchProxy{Session: &MchSession{Config: config}}, err
	}

	httpClient := http.Client{}
	res, err := httpClient.Post(
		authUrl,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return &MchProxy{Session: &MchSession{Config: config}}, err
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return &MchProxy{Session: &MchSession{Config: config}}, fmt.Errorf("invalid status code %d has been received from %s", res.StatusCode, res.Request.URL)
	}

	var token models.MchToken
	err = json.NewDecoder(res.Body).Decode(&token)

	if err != nil {
		return &MchProxy{Session: &MchSession{Config: config}}, err
	}

	session := MchSession{
		Config: config,
		Token:  &token,
	}

	_, idTokenPayload, err := utils.DecodeIdToken(token.IdToken)
	if err != nil {
		return nil, err
	}

	session.UserId = idTokenPayload.Sub

	return &MchProxy{
		HttpClient: &httpClient,
		Session:    &session,
	}, nil
}

func NewProxy(token *models.MchToken) (*MchProxy, error) {
	config, err := models.GetConfiguration()
	if err != nil {
		return nil, err
	}

	return NewProxyFromConfig(config, token), nil
}

func NewProxyFromConfig(config *models.MchConfig, token *models.MchToken) *MchProxy {
	var proxy = MchProxy{}
	proxy.HttpClient = &http.Client{}
	proxy.Session = &MchSession{}
	proxy.Session.Config = config
	proxy.Session.Token = token
	// will be set after being authenticated..
	proxy.Session.UserId = "<unknown>"
	return &proxy
}

func (mp *MchProxy) Relogin(clientId string, clientSecret string) error {

	session := mp.Session

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

	httpClient := mp.HttpClient
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
	fmt.Print(string(b))
	res.Body = io.NopCloser(bytes.NewBuffer(b))

	var token models.MchToken
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return err
	}

	// exchange old token to the new one
	mp.Session.Token = &token

	return nil
}

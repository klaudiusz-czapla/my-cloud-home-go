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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

func GetToken(clientId string, clientSecret string, username string, password string) (*MchConfig, *MchToken, error) {
	config, err := GetConfiguration()
	if err != nil {
		return nil, nil, err
	}

	authUrl := config.Data.ComponentMap.CloudServiceUrls.ServiceAuth0URL
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
		return config, nil, err
	}

	httpClient := &http.Client{}

	resp, err := httpClient.Post(
		authUrl,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return config, nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		return config, nil, fmt.Errorf("Status code %v has been received from %s", resp.StatusCode, resp.Request.URL)
	}

	respBytesArr, _ := io.ReadAll(resp.Body)
	content := string(respBytesArr)

	log.Print(content)

	var token MchToken
	err = json.Unmarshal(respBytesArr, &config)
	if err != nil {
		return config, nil, err
	}

	return config, &token, nil
}

package mch

import (
	"fmt"
	"io"
	"net/http"
)

func (mp *MchProxy) GetUserInfo() (string, error) {

	session := mp.Session
	token := session.Token

	addr := fmt.Sprintf("%s/authservice/v1/user/userinfo",
		session.Config.GetString("cloud.service.urls", "service.auth.url"))

	// which token to choose ?
	var tokenAsString = token.AccessToken
	var bearer = "Bearer " + tokenAsString

	req, _ := http.NewRequest("GET", addr, nil)
	req.Header.Add("Authorization", bearer)

	res, err := mp.HttpClient.Do(req)
	if err != nil {
		return "", nil
	}

	b, _ := io.ReadAll(res.Body)
	s := string(b)

	return s, nil
}

func (mp *MchProxy) GetUserInfoForUser(username string) (string, error) {

	session := mp.Session
	token := session.Token

	addr := fmt.Sprintf("%s/authservice/v2/auth0/user?email=%s",
		session.Config.GetString("cloud.service.urls", "service.auth.url"),
		username)

	var tokenAsString = token.AccessToken
	var bearer = "Bearer " + tokenAsString

	req, _ := http.NewRequest("GET", addr, nil)
	req.Header.Add("Authorization", bearer)

	res, err := mp.HttpClient.Do(req)
	if err != nil {
		return "", nil
	}

	b, _ := io.ReadAll(res.Body)

	return string(b), nil
}

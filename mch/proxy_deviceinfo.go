package mch

import (
	"fmt"
	"io"
	"net/http"
)

func (mp *MchProxy) GetDeviceInfoByUser(userId string) (string, error) {

	session := mp.Session
	token := session.Token

	addr := fmt.Sprintf("%s/device/v1/user/%s?pretty=true",
		session.Config.GetString("cloud.service.urls", "service.device.url"),
		userId)

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

func (mp *MchProxy) GetDeviceInfoByDevice(deviceId string) (string, error) {
	// TODO
	return "", nil
}

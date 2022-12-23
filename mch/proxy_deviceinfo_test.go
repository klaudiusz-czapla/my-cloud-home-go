package mch

import (
	"path/filepath"
	"testing"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
)

func TestGetDeviceInfoByUser(t *testing.T) {

	jsonFilePath, _ := filepath.Abs("../config.json")

	ac, err := config.NewAppConfigFromJsonFile(jsonFilePath)
	if err != nil {
		t.Error(err.Error())
	}

	proxy, err := CreateProxyForAppConfig(ac)
	if err != nil {
		t.Error(err.Error())
	}

	userid := proxy.Session.UserId

	deviceInfo, err := proxy.GetDeviceInfoByUser(userid)
	if err != nil {
		t.Error(err.Error())
	}

	if deviceInfo == "" {
		t.Error("empty device info")
	}

}

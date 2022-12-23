package mch

import (
	"path/filepath"
	"testing"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
)

func TestUserInfo(t *testing.T) {

	jsonFilePath, _ := filepath.Abs("../config.json")

	ac, err := config.NewAppConfigFromJsonFile(jsonFilePath)
	if err != nil {
		t.Error(err.Error())
	}

	proxy, err := CreateProxyForAppConfig(ac)
	if err != nil {
		t.Error(err.Error())
	}

	userInfo, err := proxy.GetUserInfo()
	if err != nil {
		t.Error(err.Error())
	}

	if userInfo == "" {
		t.Error("empty user info")
	}
}

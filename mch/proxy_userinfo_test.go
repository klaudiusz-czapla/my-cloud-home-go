package mch

import (
	"path/filepath"
	"testing"

	cmd "github.com/klaudiusz-czapla/my-cloud-home-go/cmd"
	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
)

func TestUserInfo(t *testing.T) {

	jsonFilePath, _ := filepath.Abs("../config.json")

	ac, err := config.NewAppConfigFromJsonFile(jsonFilePath)
	if err != nil {
		t.Error(err.Error())
	}

	proxy, err := cmd.CreateProxyForAppConfig(ac)
	if err != nil {
		t.Error(err.Error())
	}

	proxy
}

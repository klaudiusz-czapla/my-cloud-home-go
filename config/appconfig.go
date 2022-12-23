package config

import (
	"github.com/klaudiusz-czapla/my-cloud-home-go/utils"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DeviceName   string `json:"deviceName"`
}

func NewAppConfigFromViper(v *viper.Viper) (*AppConfig, error) {

	ac := &AppConfig{}
	if err := v.Unmarshal(ac); err != nil {
		return nil, err
	}

	return ac, nil
}

func NewAppConfigFromJsonFile(jsonFilePath string) (*AppConfig, error) {
	json, err := utils.ReadAllText(jsonFilePath)
	if err != nil {
		return nil, err
	}

	return utils.FromJson[AppConfig](json)
}

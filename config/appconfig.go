package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func NewAppConfigFromViper(v *viper.Viper) (*AppConfig, error) {

	ac := &AppConfig{}
	if err := v.Unmarshal(ac); err != nil {
		return nil, err
	}

	return ac, nil
}

package common

import "github.com/spf13/viper"

type AppConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func FromViper(v *viper.Viper) (*AppConfig, error) {

	ac := &AppConfig{}
	if err := v.Unmarshal(ac); err != nil {
		panic(err)
	}

	return ac, nil
}

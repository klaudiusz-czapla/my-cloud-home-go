package mch

import (
	"encoding/json"
	"net/http"
)

const (
	configUrl = "https://config.mycloud.com/config/v1/config"
)

func GetConfiguration() (*MchConfig, error) {
	res, err := http.Get(configUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var config MchConfig
	err = json.NewDecoder(res.Body).Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *MchConfig) GetString(section, config string) string {
	return c.Data.ComponentMap[section][config].(string)
}

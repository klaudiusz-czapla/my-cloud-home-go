package mch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MchConfig struct {
	Data struct {
		ConfigurationID string                            `json:"configurationId"`
		ComponentMap    map[string]map[string]interface{} `json:"componentMap"`
	} `json:"data"`
}

const (
	configURL = "https://config.mycloud.com/config/v1/config"
)

func GetConfiguration() (*MchConfig, error) {
	resp, err := http.Get(configURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytesArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config MchConfig
	err = json.Unmarshal(respBytesArr, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *MchConfig) GetString(section, config string) string {
	return c.Data.ComponentMap[section][config].(string)
}

package mch

import (
	"encoding/json"
	"net/http"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/serde"
)

const (
	configUrl = "https://config.mycloud.com/config/v1/config"
)

func GetConfiguration() (*serde.MchConfig, error) {
	res, err := http.Get(configUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var config serde.MchConfig
	err = json.NewDecoder(res.Body).Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

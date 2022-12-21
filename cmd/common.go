package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/klaudiusz-czapla/my-cloud-home-go/utils"
	cmd "github.com/spf13/cobra"
)

type ContextKey string

const contextProxyKey = ContextKey("proxy")

func GetProxy(context context.Context) (*mch.MchProxy, error) {
	contextProxyValue := context.Value(contextProxyKey)

	if contextProxyValue != nil {
		proxy, ok := contextProxyValue.(*mch.MchProxy)
		if !ok {
			return nil, fmt.Errorf("invalid type: pointer to proxy object expected")
		}

		return proxy, nil
	}

	return nil, nil
}

func CreateProxyForAppConfig(ac *common.AppConfig) (*mch.MchProxy, error) {
	p, err := mch.Login(ac.ClientId, ac.ClientSecret, ac.Username, ac.Password)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetOrCreateProxy(cmd *cmd.Command, ac *common.AppConfig) (*mch.MchProxy, error) {
	proxy, err := GetProxy(cmd.Context())

	if proxy != nil {
		return proxy, nil
	}

	if err != nil {
		return nil, err
	}

	return CreateProxyForAppConfig(ac)
}

func CreateProxyForToken(ac *common.AppConfig, tokenFilePath string, token string) (*mch.MchProxy, error) {

	var tokenString = ""

	if tokenFilePath != "" {
		t, err := utils.ReadAllText(tokenFilePath)
		if err != nil {
			log.Fatal(err.Error())
		}
		tokenString = t
	} else if token != "" {
		tokenString = token
	} else {
		log.Fatalf("token file path and token cannot be both empty")
	}

	var mt mch.MchToken
	err := json.NewDecoder(strings.NewReader(tokenString)).Decode(&mt)
	if err != nil {
		return nil, err
	}

	return mch.NewProxy(&mt)
}

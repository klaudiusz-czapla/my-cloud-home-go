package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
	cmd "github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ContextKey string

const contextProxyKey = ContextKey("proxy")

func GetProxy(cmd *cmd.Command) (*mch.MchProxy, error) {
	contextProxyValue := cmd.Context().Value(contextProxyKey)

	if contextProxyValue != nil {
		proxy, ok := contextProxyValue.(*mch.MchProxy)
		if !ok {
			return nil, fmt.Errorf("invalid type: MchProxy pointer expected")
		}

		return proxy, nil
	}

	return nil, nil
}

func CreateProxy(v *viper.Viper) (*mch.MchProxy, error) {
	p, err := mch.Login(v.GetString("clientId"), v.GetString("clientSecret"), v.GetString("username"), v.GetString("password"))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetOrCreateProxy(cmd *cmd.Command, v *viper.Viper) (*mch.MchProxy, error) {
	proxy, err := GetProxy(cmd)

	if proxy != nil {
		return proxy, nil
	}

	if err != nil {
		return nil, err
	}

	return CreateProxy(v)
}

func CreateProxyEitherFromPlainTextOrFile(ac *AppConfig) (*mch.MchProxy, error) {
	var clientId = v.GetString("clientId")
	var clientSecret = v.GetString("clientSecret")
	var t = v.GetString("token")
	var f = v.GetString("from")

	var tokenString = ""

	if f != "" {
		t, err := utils.ReadFileContent(f)
		if err != nil {
			log.Fatal(err.Error())
		}
		tokenString = t
	} else if t != "" {
		tokenString = t
	} else {
		log.Fatal("unknown command switch")
	}

	var token mch.MchToken
	err := json.NewDecoder(strings.NewReader(tokenString)).Decode(&token)
	if err != nil {
		return err
	}
	proxy, err := mch.NewProxy(&token)
	if err != nil {
		log.Fatal(err.Error())
	}
}

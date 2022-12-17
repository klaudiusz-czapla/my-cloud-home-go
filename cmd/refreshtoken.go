package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitRefreshTokenCommand(v *viper.Viper) *cobra.Command {
	var refreshTokenCmd = &cobra.Command{
		Use:   "refresh-token",
		Short: "Refresh token",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			var clientId = v.GetString("clientId")
			var clientSecret = v.GetString("clientSecret")
			var tokenString = v.GetString("token")

			config, err := mch.GetConfiguration()
			if err != nil {
				log.Fatal(err.Error())
			}

			var proxy = mch.MchProxy{}
			proxy.HttpClient = &http.Client{}
			proxy.MchSession = &mch.MchSession{}
			proxy.Config = config
			var token mch.MchToken
			err = json.NewDecoder(strings.NewReader(tokenString)).Decode(&token)
			proxy.MchSession.Token = &token

			err = proxy.Relogin(clientId, clientSecret)
			if err != nil {
				log.Fatal(err.Error())
			}
		},
	}

	refreshTokenCmd.Flags().StringP("token", "t", "", "Refresh token from the original one (which is about to expire soon).")
	v.BindPFlag("token", refreshTokenCmd.Flags().Lookup("token"))

	return refreshTokenCmd
}

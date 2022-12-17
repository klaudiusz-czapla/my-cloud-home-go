package cmd

import (
	"encoding/json"
	"log"
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
		PreRun: func(cmd *cobra.Command, args []string) {
			//var clientId = v.GetString("clientId")
			//log.Printf("Client id=%s", clientId)

			//var clientSecret = v.GetString("clientSecret")
			//log.Printf("Client secret=%s", clientSecret)

			//var token = v.GetString("token")
			//log.Printf("Token=%s", token)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var clientId = v.GetString("clientId")
			var clientSecret = v.GetString("clientSecret")
			var tokenString = v.GetString("token")

			var token mch.MchToken
			err := json.NewDecoder(strings.NewReader(tokenString)).Decode(&token)
			if err != nil {
				log.Fatal(err.Error())
			}

			proxy, err := mch.NewProxy(&token)
			if err != nil {
				log.Fatal(err.Error())
			}

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

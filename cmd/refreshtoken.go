package cmd

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const refreshTokenCmdName = "refresh-token"

func InitRefreshTokenCommand(v *viper.Viper) *cobra.Command {
	var refreshTokenCmd = &cobra.Command{
		Use:   refreshTokenCmdName,
		Short: "Refresh token",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", refreshTokenCmdName)
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
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", refreshTokenCmdName)
		},
	}

	refreshTokenCmd.Flags().StringP("token", "t", "", "Token.")
	refreshTokenCmd.Flags().StringP("from", "f", "", "Token file")

	v.BindPFlag("token", refreshTokenCmd.Flags().Lookup("token"))
	v.BindPFlag("from", refreshTokenCmd.Flags().Lookup("from"))

	return refreshTokenCmd
}

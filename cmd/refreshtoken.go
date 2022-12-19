package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
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
			var t = v.GetString("token")
			var f = v.GetString("from")

			var tokenString = ""

			if f != "" {
				if !utils.FileExists(f) {
					log.Fatal("file does not exist")
				}
				data, err := os.ReadFile(f)
				if err != nil {
					log.Fatal(err.Error())
				}
				tokenString = string(data)
			} else if t != "" {
				tokenString = t
			} else {
				log.Fatal("unknown command switch")
			}

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

	refreshTokenCmd.Flags().String("token", "", "Token.")
	refreshTokenCmd.Flags().String("from", "", "Token file")
	refreshTokenCmd.MarkFlagsMutuallyExclusive("token", "from")

	v.BindPFlag("token", refreshTokenCmd.Flags().Lookup("token"))
	v.BindPFlag("from", refreshTokenCmd.Flags().Lookup("from"))

	return refreshTokenCmd
}

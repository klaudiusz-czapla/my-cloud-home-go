package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const REFRESH_TOKEN_CMD_NAME = "refresh-token"

const (
	REFRESH_TOKEN_CMD_TOKEN_FLAG_LBL = "token"
	REFRESH_TOKEN_CMD_FROM_FLAG_LBL  = "from"
)

func InitRefreshTokenCommand(v *viper.Viper) *cobra.Command {
	var refreshTokenCmd = &cobra.Command{
		Use:   REFRESH_TOKEN_CMD_NAME,
		Short: "Refresh token",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", REFRESH_TOKEN_CMD_NAME)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var clientId = v.GetString(clientIdLabel)
			var clientSecret = v.GetString(clientSecretLabel)
			var t = v.GetString("token")
			var f = v.GetString("from")

			proxy, err := CreateProxyEitherFromPlainTextOrFile(v)

			err = proxy.Relogin(clientId, clientSecret)
			if err != nil {
				log.Fatal(err.Error())
			}

			tokenAsBytes, _ := json.Marshal(proxy.Session.Token)
			tokenAsString := string(tokenAsBytes)

			if f != "" {
				err := utils.WriteFileContent(f, tokenString)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else if t != "" {
				// if token received from terminal then print it to terminal as well instead of file
				fmt.Print(tokenAsString)
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

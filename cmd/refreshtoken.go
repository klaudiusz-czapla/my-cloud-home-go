package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
	"github.com/klaudiusz-czapla/my-cloud-home-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const refreshTokenCmdName = "refresh-token"

const (
	refreshCmdTokenFlag = "token"
	refreshCmdFromFlag  = "from"
)

func InitRefreshTokenCommand(v *viper.Viper) *cobra.Command {

	ac, err := config.NewAppConfigFromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var refreshTokenCmd = &cobra.Command{
		Use:   refreshTokenCmdName,
		Short: "Refresh token",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", refreshTokenCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var tokenFilePath = v.GetString(refreshTokenCmdName + "." + refreshCmdFromFlag)
			var token = v.GetString(refreshTokenCmdName + "." + refreshCmdTokenFlag)

			proxy, err := CreateProxyForToken(ac, tokenFilePath, token)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = proxy.Relogin(ac.ClientId, ac.ClientSecret)
			if err != nil {
				log.Fatal(err.Error())
			}

			tokenAsBytes, _ := json.Marshal(proxy.Session.Token)
			tokenAsString := string(tokenAsBytes)

			if tokenFilePath != "" {
				err := utils.WriteAllText(tokenFilePath, tokenAsString)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else if token != "" {
				// if token received from terminal then print it to terminal as well instead of file
				fmt.Print(tokenAsString)
			} else {
				log.Fatalf("token file path and token cannot be both empty. Either the first one or the second parameter has to be set to some non-empty value")
			}

		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", refreshTokenCmdName)
		},
	}

	refreshTokenCmd.Flags().String(refreshCmdTokenFlag, "", "Token")
	refreshTokenCmd.Flags().String(refreshCmdFromFlag, "", "Token file")
	refreshTokenCmd.MarkFlagsMutuallyExclusive(refreshCmdTokenFlag, refreshCmdFromFlag)

	v.BindPFlag(refreshTokenCmdName+"."+refreshCmdTokenFlag, refreshTokenCmd.Flags().Lookup(refreshCmdTokenFlag))
	v.BindPFlag(refreshTokenCmdName+"."+refreshCmdFromFlag, refreshTokenCmd.Flags().Lookup(refreshCmdFromFlag))

	return refreshTokenCmd
}

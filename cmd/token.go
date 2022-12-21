package cmd

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/klaudiusz-czapla/my-cloud-home-go/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tokenCmdName = "token"

const (
	tokenCmdToFlag = "to"
)

func InitTokenCommand(v *viper.Viper) *cobra.Command {

	ac, err := common.NewAppConfigFromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var tokenCmd = &cobra.Command{
		Use:              tokenCmdName,
		Short:            "Get the user token",
		Long:             ``,
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			proxy, err := CreateProxyForAppConfig(ac)
			if err != nil {
				log.Fatal(err.Error())
			}

			cmd.SetContext(context.WithValue(cmd.Context(), contextProxyKey, proxy))
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", tokenCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var tokenFilePath = v.GetString(tokenCmdName + "." + tokenCmdToFlag)

			proxy, err := GetProxyFromContext(cmd.Context())
			if err != nil {
				log.Fatal(err.Error())
			}

			json.NewEncoder(os.Stdout).Encode(proxy.Session.Token)

			if tokenFilePath != "" {

				tokenAsBytes, _ := json.Marshal(proxy.Session.Token)
				tokenAsString := string(tokenAsBytes)
				err := utils.WriteAllText(tokenFilePath, tokenAsString)
				if err != nil {
					log.Fatal(err.Error())
				}

			}
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", tokenCmdName)
		},
	}

	tokenCmd.Flags().String(tokenCmdToFlag, "", "Token file")

	v.BindPFlag(tokenCmdName+"."+tokenCmdToFlag, tokenCmd.Flags().Lookup(tokenCmdToFlag))

	return tokenCmd
}

package cmd

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tokenCmdName = "token"

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
			proxy, err := GetProxy(cmd)
			if err != nil {
				log.Fatal(err.Error())
			}

			json.NewEncoder(os.Stdout).Encode(proxy.Session.Token)

			if v.GetString("to") != "" {

				file, err := os.OpenFile(v.GetString("to"), os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(int(0600)))
				if err != nil {
					log.Fatal(err.Error())
				}

				tokenAsBytes, _ := json.Marshal(proxy.Session.Token)
				tokenAsString := string(tokenAsBytes)
				file.WriteString(tokenAsString)

				if err := file.Close(); err != nil {
					log.Fatal(err.Error())
				}
			}
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", tokenCmdName)
		},
	}

	tokenCmd.Flags().String("to", "", "Token file")

	v.BindPFlag("to", tokenCmd.Flags().Lookup("to"))

	return tokenCmd
}

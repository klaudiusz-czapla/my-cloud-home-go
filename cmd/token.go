package cmd

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tokenCmdName = "token"

func InitTokenCommand(v *viper.Viper) *cobra.Command {
	var tokenCmd = &cobra.Command{
		Use:              tokenCmdName,
		Short:            "Get the user token",
		Long:             ``,
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			proxy, err := mch.Login(v.GetString("clientId"), v.GetString("clientSecret"), v.GetString("username"), v.GetString("password"))
			if err != nil {
				log.Fatal(err.Error())
			}

			cmd.SetContext(context.WithValue(cmd.Context(), contextProxyKey, proxy))
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", tokenCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {
			contextProxyValue := cmd.Context().Value(contextProxyKey)
			if contextProxyValue == nil {
				log.Fatal("empty proxy object received from context")
			}
			proxy := contextProxyValue.(mch.MchProxy)

			json.NewEncoder(os.Stdout).Encode(proxy.Session.Token)

			if v.GetString("to") != "" {

				file, err := os.OpenFile(v.GetString("as"), os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(int(0600)))
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
	}

	tokenCmd.Flags().String("to", "", "Token file")

	v.BindPFlag("to", tokenCmd.Flags().Lookup("to"))

	return tokenCmd
}

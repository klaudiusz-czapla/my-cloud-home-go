package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitTokenCommand(v *viper.Viper) *cobra.Command {
	var tokenCmd = &cobra.Command{
		Use:   "token",
		Short: "Get the user token",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			session, err := mch.Login(v.GetString("clientId"), v.GetString("clientSecret"), v.GetString("username"), v.GetString("password"))
			if err != nil {
				log.Fatal(err.Error())
			}

			propName := v.GetString("propertyName")
			if propName == "" {
				json.NewEncoder(os.Stdout).Encode(session.Token)
				return
			}

			switch propName {
			case "id_token":
				fmt.Print(args)
			case "refresh_token":
				fmt.Print(args)
			default:
				return
			}
		},
	}

	tokenCmd.Flags().String("propertyName", "", "Json property name meant to be extracted.")
	v.BindPFlag("propertyName", tokenCmd.Flags().Lookup("propertyName"))

	return tokenCmd
}

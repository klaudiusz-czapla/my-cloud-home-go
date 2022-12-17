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
			proxy, err := mch.Login(v.GetString("clientId"), v.GetString("clientSecret"), v.GetString("username"), v.GetString("password"))
			if err != nil {
				log.Fatal(err.Error())
			}

			json.NewEncoder(os.Stdout).Encode(proxy.Session.Token)

			if v.GetString("as") != "" {

				file, err := os.OpenFile(v.GetString("as"), os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(int(0600)))
				if err != nil {
					log.Fatal(err.Error())
				}

				if err := file.Close(); err != nil {
					log.Fatal(err.Error())
				}
			}

			if v.GetBool("decode-id-token") {
				claims, _ := mch.DecodeToken(proxy.Session.Token.IdToken)
				fmt.Println(claims)
			}

			if v.GetBool("decode-refresh-token") {

			}
		},
	}

	tokenCmd.Flags().String("as", "", "Save token as file")
	tokenCmd.Flags().Bool("decode-id-token", false, "Decode id token.")
	tokenCmd.Flags().Bool("decode-access-token", false, "Decode access token.")

	v.BindPFlag("as", tokenCmd.Flags().Lookup("as"))
	v.BindPFlag("decode-id-token", tokenCmd.Flags().Lookup("decode-id-token"))
	v.BindPFlag("decode-refresh-token", tokenCmd.Flags().Lookup("decode-refresh-token"))

	return tokenCmd
}

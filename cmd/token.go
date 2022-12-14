package cmd

import (
	"encoding/json"
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

			json.NewEncoder(os.Stdout).Encode(session.Token)

			// TODO: move it to separate command
			// keep possibility to persist token (save it to the file for later usage)
			if v.GetBool("refresh") {
				err = mch.Relogin(v.GetString("clientId"), v.GetString("clientSecret"), session)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		},
	}

	tokenCmd.Flags().BoolP("refresh", "r", false, "Refresh the token.")
	v.BindPFlag("refresh", tokenCmd.Flags().Lookup("refresh"))

	return tokenCmd
}

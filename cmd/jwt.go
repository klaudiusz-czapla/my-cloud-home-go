package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const jwtCmdName = "jwt"

const ()

// jwt command can be subcommand for token (parent) command
// at the same time it can be standalone command - in that case all input parameters will be processed directly here
// instead of relying on getting context with expected value from parent command
func InitJwtCommand(v *viper.Viper) *cobra.Command {

	ac, err := common.FromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var jwtCmd = &cobra.Command{
		Use:   jwtCmdName,
		Short: "Aggregates operations which can be performed on JWT token",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", jwtCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var proxy *mch.MchProxy

			var parentCmd = cmd.Parent()
			// executed as subcommand of token parent command
			if parentCmd.Use == "token" {
				proxy, err := GetProxy(cmd.Context())
				if err != nil {
					log.Fatal(err.Error())
				}

			} else {
				CreateProxyForToken(ac, v.GetString(), v.GetString())
			}

			if proxy == nil {
				log.Fatal("empty proxy object")
			}

			if v.GetBool("decode-id-token") {
				claims, _ := mch.DecodeToken(proxy.Session.Token.IdToken)
				fmt.Print(claims)
			}

			if v.GetBool("decode-access-token") {
				claims, _ := mch.DecodeToken(proxy.Session.Token.AccessToken)
				fmt.Print(claims)
			}

		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", jwtCmdName)
		},
	}

	jwtCmd.Flags().Bool("decode-id-token", false, "Decode id token.")
	jwtCmd.Flags().Bool("decode-access-token", false, "Decode access token.")

	v.BindPFlag("decode-id-token", jwtCmd.Flags().Lookup("decode-id-token"))
	v.BindPFlag("decode-access-token", jwtCmd.Flags().Lookup("decode-access-token"))

	jwtCmd.Flags().String("token", "", "Token.")
	jwtCmd.Flags().String("from", "", "Token file")
	jwtCmd.MarkFlagsMutuallyExclusive("token", "from")

	v.BindPFlag("token", jwtCmd.Flags().Lookup("token"))
	v.BindPFlag("from", jwtCmd.Flags().Lookup("from"))

	return jwtCmd
}

package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const jwtCmdName = "jwt"

const (
	jwtCmdTokenFlag             = "token"
	jwtCmdFromFlag              = "from"
	jwtCmdDecodeIdTokenFlag     = "decode-id-token"
	jwtCmdDecodeAccessTokenFlag = "decode-access-token"
)

// jwt command can be subcommand for token (parent) command
// at the same time it can be standalone command - in that case all input parameters will be processed directly here
// instead of relying on getting context with expected value from parent command
func InitJwtCommand(v *viper.Viper) *cobra.Command {

	ac, err := common.NewAppConfigFromViper(v)
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
				proxy, err = GetProxyFromContext(cmd.Context())
				if err != nil {
					log.Fatal(err.Error())
				}

			} else {
				CreateProxyForToken(ac, v.GetString(jwtCmdName+"."+jwtCmdTokenFlag), v.GetString(jwtCmdName+"."+jwtCmdFromFlag))
			}

			if proxy == nil {
				log.Fatal("empty proxy object")
			}

			if v.GetBool(jwtCmdName + "." + jwtCmdDecodeIdTokenFlag) {
				claims, _, _ := utils.DecodeToken(proxy.Session.Token.IdToken)
				fmt.Print(claims)
			}

			if v.GetBool(jwtCmdName + "." + jwtCmdDecodeAccessTokenFlag) {
				claims, _, _ := utils.DecodeToken(proxy.Session.Token.AccessToken)
				fmt.Print(claims)
			}

		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", jwtCmdName)
		},
	}

	jwtCmd.Flags().Bool(jwtCmdDecodeIdTokenFlag, false, "Decode id token.")
	jwtCmd.Flags().Bool(jwtCmdDecodeAccessTokenFlag, false, "Decode access token.")
	jwtCmd.Flags().String(jwtCmdTokenFlag, "", "Token.")
	jwtCmd.Flags().String(jwtCmdFromFlag, "", "Token file")
	jwtCmd.MarkFlagsMutuallyExclusive(jwtCmdTokenFlag, jwtCmdFromFlag)

	v.BindPFlag(jwtCmdName+"."+jwtCmdDecodeIdTokenFlag, jwtCmd.Flags().Lookup(jwtCmdDecodeIdTokenFlag))
	v.BindPFlag(jwtCmdName+"."+jwtCmdDecodeAccessTokenFlag, jwtCmd.Flags().Lookup(jwtCmdDecodeAccessTokenFlag))
	v.BindPFlag(jwtCmdName+"."+jwtCmdTokenFlag, jwtCmd.Flags().Lookup(jwtCmdTokenFlag))
	v.BindPFlag(jwtCmdName+"."+jwtCmdFromFlag, jwtCmd.Flags().Lookup(jwtCmdFromFlag))

	return jwtCmd
}

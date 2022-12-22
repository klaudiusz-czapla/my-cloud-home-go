package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
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
func InitJwtCommand(v *viper.Viper, parent *cobra.Command) *cobra.Command {

	buildFlagName := func(f string) string {
		if parent == nil {
			return jwtCmdName + "." + f
		} else {
			return parent.Use + "." + jwtCmdName + "." + f
		}
	}

	ac, err := config.NewAppConfigFromViper(v)
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
				proxy, err = CreateProxyForToken(
					ac,
					v.GetString(buildFlagName(jwtCmdFromFlag)),
					v.GetString(buildFlagName(jwtCmdTokenFlag)))
				if err != nil {
					log.Fatal(err.Error())
				}
			}

			if proxy == nil {
				log.Fatal("empty proxy object")
			}

			if v.GetBool(buildFlagName(jwtCmdDecodeIdTokenFlag)) {
				claims, _, _ := utils.DecodeIdToken(proxy.Session.Token.IdToken)
				fmt.Print(claims)
			}

			if v.GetBool(buildFlagName(jwtCmdDecodeAccessTokenFlag)) {
				claims, _, _ := utils.DecodeAccessToken(proxy.Session.Token.AccessToken)
				fmt.Print(claims)
			}

		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", jwtCmdName)
		},
	}

	jwtCmd.Flags().Bool(jwtCmdDecodeIdTokenFlag, false, "Decode id token.")
	jwtCmd.Flags().Bool(jwtCmdDecodeAccessTokenFlag, false, "Decode access token.")
	jwtCmd.Flags().String(jwtCmdTokenFlag, "", "Token")
	jwtCmd.Flags().String(jwtCmdFromFlag, "", "Token file")
	jwtCmd.MarkFlagsMutuallyExclusive(jwtCmdTokenFlag, jwtCmdFromFlag)

	v.BindPFlag(buildFlagName(jwtCmdDecodeIdTokenFlag), jwtCmd.Flags().Lookup(jwtCmdDecodeIdTokenFlag))
	v.BindPFlag(buildFlagName(jwtCmdDecodeAccessTokenFlag), jwtCmd.Flags().Lookup(jwtCmdDecodeAccessTokenFlag))
	v.BindPFlag(buildFlagName(jwtCmdTokenFlag), jwtCmd.Flags().Lookup(jwtCmdTokenFlag))
	v.BindPFlag(buildFlagName(jwtCmdFromFlag), jwtCmd.Flags().Lookup(jwtCmdFromFlag))

	return jwtCmd
}

package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const userInfoCmdName = "user-info"

const (
	userInfoCmdFromFlag = "from"
)

func InitUserInfoCommand(v *viper.Viper) *cobra.Command {

	ac, err := common.NewAppConfigFromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var userInfoCmd = &cobra.Command{
		Use:   userInfoCmdName,
		Short: "Retrieves user info in json format",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", userInfoCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var tokenFilePath = v.GetString(userInfoCmdName + "." + userInfoCmdFromFlag)

			proxy, err := CreateProxyForToken(ac, tokenFilePath, "")
			if err != nil {
				log.Fatal(err.Error())
			}

			userinfo, err := proxy.GetUserInfo()
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Print(userinfo)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", userInfoCmdName)
		},
	}

	userInfoCmd.Flags().String(refreshCmdFromFlag, "", "Token file")

	v.BindPFlag(userInfoCmdName+"."+userInfoCmdFromFlag, userInfoCmd.Flags().Lookup(userInfoCmdFromFlag))

	return userInfoCmd
}

package cmd

import (
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/common"
	"github.com/spf13/cobra"
)

const userInfoCmdName = "user-info"

func InitUserInfoCommand() *cobra.Command {

	ac, err := common.NewAppConfigFromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var userInfoCmd = &cobra.Command{
		Use:   userInfoCmdName,
		Short: "Retrieves user info in json format",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", configCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {
			proxy, err := CreateProxyForAppConfig(ac)
			if err != nil {
				log.Fatal(err.Error())
			}
			_ = proxy.GetUserInfo(ac.Username)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", configCmdName)
		},
	}

	return userInfoCmd
}

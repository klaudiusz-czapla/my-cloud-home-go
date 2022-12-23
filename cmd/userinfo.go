package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const userInfoCmdName = "user-info"

const (
	userInfoCmdFromFlag    = "from"
	userInfoByUserNameFlag = "byUserName"
)

func InitUserInfoCommand(v *viper.Viper) *cobra.Command {

	ac, err := config.NewAppConfigFromViper(v)
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
			var byUserName = v.GetBool(userInfoCmdName + "." + userInfoByUserNameFlag)

			proxy, err := mch.CreateProxyForToken(ac, tokenFilePath, "")
			if err != nil {
				log.Fatal(err.Error())
			}

			if byUserName {
				userInfo, err := proxy.GetUserInfoByUserName(ac.Username)
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Print(userInfo)
			} else {
				userInfo, err := proxy.GetUserInfo()
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Print(userInfo)
			}
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", userInfoCmdName)
		},
	}

	userInfoCmd.Flags().String(userInfoCmdFromFlag, "", "Token file")
	userInfoCmd.Flags().Bool(userInfoByUserNameFlag, false, "By user name")

	v.BindPFlag(userInfoCmdName+"."+userInfoCmdFromFlag, userInfoCmd.Flags().Lookup(userInfoCmdFromFlag))
	v.BindPFlag(userInfoCmdName+"."+userInfoByUserNameFlag, userInfoCmd.Flags().Lookup(userInfoByUserNameFlag))

	return userInfoCmd
}

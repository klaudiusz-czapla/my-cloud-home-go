package cmd

import (
	"fmt"
	"log"

	"github.com/klaudiusz-czapla/my-cloud-home-go/config"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const deviceInfoCmdName = "device-info"

const (
	deviceInfoCmdFromFlag = "from"
)

func InitDeviceInfoCommand(v *viper.Viper) *cobra.Command {

	ac, err := config.NewAppConfigFromViper(v)
	if err != nil {
		log.Fatal(err.Error())
	}

	var deviceInfoCmd = &cobra.Command{
		Use:   deviceInfoCmdName,
		Short: "Retrieves device info for currently logged-in user",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Printf("executing '%s' command..", deviceInfoCmdName)
		},
		Run: func(cmd *cobra.Command, args []string) {

			var tokenFilePath = v.GetString(deviceInfoCmdName + "." + deviceInfoCmdFromFlag)

			proxy, err := mch.CreateProxyForToken(ac, tokenFilePath, "")
			if err != nil {
				log.Fatal(err.Error())
			}

			userid := proxy.Session.UserId
			deviceInfo, err := proxy.GetDeviceInfoByUser(userid)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Print(deviceInfo)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Printf("command '%s' has been executed..", userInfoCmdName)
		},
	}

	deviceInfoCmd.Flags().String(deviceInfoCmdFromFlag, "", "Token file")

	v.BindPFlag(deviceInfoCmdName+"."+deviceInfoCmdFromFlag, deviceInfoCmd.Flags().Lookup(deviceInfoCmdFromFlag))

	return deviceInfoCmd
}

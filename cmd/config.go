package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/models"
	"github.com/spf13/cobra"
)

const configCmdName = "config"

var configCmd = &cobra.Command{
	Use:   configCmdName,
	Short: "Retrieves configuration in json format",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Printf("executing '%s' command..", configCmdName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := models.GetConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}

		json.NewEncoder(os.Stdout).Encode(c)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Printf("command '%s' has been executed..", configCmdName)
	},
}

func InitConfigCommand() *cobra.Command {
	return configCmd
}

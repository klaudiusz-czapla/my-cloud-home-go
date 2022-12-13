package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Retrieves configuration in json format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := mch.GetConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}

		json.NewEncoder(os.Stdout).Encode(c)
	},
}

func InitConfigCommand() *cobra.Command {
	return configCmd
}

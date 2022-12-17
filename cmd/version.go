package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const (
	versionCmdName = "version"
	version        = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   versionCmdName,
	Short: "Print the version number of My Cloud Home CLI app",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Printf("executing '%s' command..", configCmdName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Printf("command '%s' has been executed..", versionCmdName)
	},
}

func InitVersionCommand() *cobra.Command {
	// nothing to be initialized
	return versionCmd
}

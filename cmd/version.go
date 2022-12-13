package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of My Cloud Home CLI app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func InitVersionCommand() *cobra.Command {
	// nothing to be initialized
	return versionCmd
}

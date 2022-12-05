package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	version = "0.0.1"
)

var rootCmd = &cobra.Command{
	Use:   "my-cloud-home-go",
	Short: "My Cloud Home CLI application",
	Long:  `my-cloud-home-go is meant for managing My Cloud Home devices`,
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of My Cloud Home CLI app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func main() {

	rootCmd.PersistentFlags().StringP("username", "u", "", "WD My Cloud Home user name")
	rootCmd.PersistentFlags().StringP("password", "p", "", "WD My Cloud Home user password")
	rootCmd.MarkFlagsRequiredTogether("username", "password")

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.Execute()
}

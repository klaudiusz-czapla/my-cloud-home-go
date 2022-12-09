package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	Short: "Get the user token",
	Long:  ``,
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

	path := os.Args[0]
	absPath, _ := filepath.Abs(path)

	log.Print("App has been started..")
	log.Printf("Started in %s", absPath)

	rootCmd.PersistentFlags().StringP("username", "u", "", "WD My Cloud Home user name")
	rootCmd.PersistentFlags().StringP("password", "p", "", "WD My Cloud Home user password")
	rootCmd.PersistentFlags().StringP("clientid", "c", "", "Client Id")
	rootCmd.PersistentFlags().StringP("clientsecret", "s", "", "Client Secret")

	viper.SetConfigName("mch")
	viper.SetConfigType("ini")
	viper.AddConfigPath(path)
	viper.AddConfigPath("~/.mch")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mch")
	viper.AutomaticEnv()

	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("clientid", rootCmd.PersistentFlags().Lookup("clientid"))
	viper.BindPFlag("clientsecret", rootCmd.PersistentFlags().Lookup("clientsecret"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.Execute()
}

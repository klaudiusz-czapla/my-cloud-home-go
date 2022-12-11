package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	version = "0.0.1"
)

var (
	configPath   string
	username     string
	password     string
	clientId     string
	clientSecret string
)

var rootCmd = &cobra.Command{
	Use:   "my-cloud-home-go",
	Short: "My Cloud Home CLI application",
	Long:  `my-cloud-home-go is meant for managing My Cloud Home devices`,
	PreRun: func(cmd *cobra.Command, args []string) {
		print("asd")
	},
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

	//path := os.Args[0]
	absPath, _ := filepath.Abs(".")

	cobra.OnInitialize(func() {
		if configPath != "" {
			if mch.FileExists(configPath) {
				configFileExt := filepath.Ext(configPath)

				if configFileExt == "ini" {
					viper.AddConfigPath(configPath)
					viper.ReadInConfig()
				}
			}
		}
	})

	log.Print("App has been started..")
	log.Printf("Started from the path: %s", absPath)

	rootCmd.Flags().StringVar(&configPath, "configPath", "", "Configuration file path.")

	tokenCmd.PersistentFlags().StringVar(&username, "username", "", "WD My Cloud Home user name.")
	tokenCmd.PersistentFlags().StringVar(&password, "password", "", "WD My Cloud Home user password")
	tokenCmd.PersistentFlags().StringVar(&clientId, "clientId", "", "Client Id")
	tokenCmd.PersistentFlags().StringVar(&clientSecret, "clientSecret", "", "Client Secret")

	viper.SetConfigName("mch")
	viper.SetConfigType("ini")
	viper.AddConfigPath(absPath)
	viper.AddConfigPath("~/.mch")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mch")
	viper.AutomaticEnv()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tokenCmd)

	viper.Debug()

	rootCmd.Execute()
}

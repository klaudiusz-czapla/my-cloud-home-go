package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

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
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get the user token",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.Debug()
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, token, err := mch.GetToken(clientId, clientSecret, username, password)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(token)
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

				if configFileExt == ".ini" {
					viper.AddConfigPath(configPath)
					var err = viper.ReadInConfig()
					if err != nil {
						log.Fatal(err.Error())
					}
				}
			}
		}
	})

	log.Print("App has been started..")
	log.Printf("Started from the path: %s", absPath)

	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "Configuration file path.")
	viper.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("configPath"))

	tokenCmd.Flags().StringVar(&username, "username", "", "WD My Cloud Home user name.")
	tokenCmd.Flags().StringVar(&password, "password", "", "WD My Cloud Home user password")
	tokenCmd.Flags().StringVar(&clientId, "clientId", "", "Client Id")
	tokenCmd.Flags().StringVar(&clientSecret, "clientSecret", "", "Client Secret")
	viper.BindPFlag("username", tokenCmd.Flags().Lookup("username"))
	viper.BindPFlag("password", tokenCmd.Flags().Lookup("password"))
	viper.BindPFlag("clientId", tokenCmd.Flags().Lookup("clientId"))
	viper.BindPFlag("clientSecret", tokenCmd.Flags().Lookup("clientSecret"))

	viper.AddConfigPath(absPath)
	viper.SetConfigType("ini")
	viper.SetConfigName("mch")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	viper.SetEnvPrefix("mch")
	viper.AutomaticEnv()

	// debug mode
	if viper.GetString(strings.ToUpper("clientId")) == "" {
		log.Fatal("ClientId has empty value")
	}

	// debug mode
	if viper.GetString(strings.ToUpper("clientSecret")) == "" {
		log.Fatal("ClientSecret has empty value")
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.Execute()
}

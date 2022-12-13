package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	version = "0.0.1"
)

var (
	//home       = os.Getenv("HOME")
	absPath, _ = filepath.Abs(".")
	v          = viper.New()
)

const (
	defaultConfigFileName string = "config"
)

var (
	defaultConfigPath string = absPath
)

var (
	configFileName string = defaultConfigFileName
	configPath     string = defaultConfigPath
)

var rootCmd = &cobra.Command{
	Use:   "my-cloud-home-go",
	Short: "My Cloud Home CLI application",
	Long:  `my-cloud-home-go is meant for managing My Cloud Home devices`,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Retrieves configuration in json format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := mch.GetConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}

		json.NewEncoder(log.Writer()).Encode(c)
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get the user token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		session, err := mch.Login(v.GetString("clientId"), v.GetString("clientSecret"), v.GetString("username"), v.GetString("password"))
		if err != nil {
			log.Fatal(err.Error())
		}

		json.NewEncoder(log.Writer()).Encode(session.Token)
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

func addJsonFile(filePath string, fileName string) {
	const ext = "json"
	in := filePath + string(os.PathSeparator) + fileName + "." + ext
	if in != "" {
		if mch.FileExists(in) {
			v.AddConfigPath(filePath)
			v.SetConfigType(ext)
			v.SetConfigFile(in)
			var err = v.ReadInConfig()
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Printf("configuration has been retrieved from file: %s", in)
		}
	}
}

func main() {

	log.Print("app has been started..")
	log.Printf("started from the path: %s", absPath)

	v.SetDefault("configFileName", defaultConfigFileName)
	v.SetDefault("configPath", defaultConfigPath)

	rootCmd.PersistentFlags().StringVarP(&configFileName, "configFileName", "c", defaultConfigFileName, "Configuration file name.")
	rootCmd.PersistentFlags().StringVarP(&configPath, "configPath", "p", defaultConfigPath, "Configuration path.")
	v.BindPFlag("configFileName", rootCmd.PersistentFlags().Lookup("configFileName"))
	v.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("configPath"))

	if cp := v.GetString("configPath"); cp != "" {
		if cf := v.GetString("configFileName"); cf != "" {
			addJsonFile(cp, cf)
		}
	}

	cobra.OnInitialize(func() {
		if configPath != defaultConfigPath || configFileName != defaultConfigFileName {
			addJsonFile(configPath, configFileName)
		}
	})

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.Execute()
}

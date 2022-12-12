package main

import (
	"fmt"
	"log"
	"os"
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
	//home       = os.Getenv("HOME")
	absPath, _ = filepath.Abs(".")
	v          = viper.New()
)

const (
	defaultConfigFileName string = "config"
)

var (
	configFileName    string = defaultConfigFileName
	defaultConfigPath string = absPath
	configPath        string = defaultConfigPath
	username          string
	password          string
	clientId          string
	clientSecret      string
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
		mch.BindFlags(cmd, v)
	},
	Run: func(cmd *cobra.Command, args []string) {
		un := viper.GetString("username")
		_, token, err := mch.GetToken(clientId, clientSecret, un, password)
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

func addJsonFile(filePath string, fileName string) {
	const ext = "json"
	in := filePath + string(os.PathSeparator) + fileName + "." + ext
	if in != "" {
		if mch.FileExists(in) {
			v.AddConfigPath(filePath)
			v.SetConfigType("json")
			v.SetConfigName(fileName)
			var err = v.ReadInConfig()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func main() {

	log.Print("App has been started..")
	log.Printf("Started from the path: %s", absPath)

	v.SetDefault("configFileName", defaultConfigFileName)
	v.SetDefault("configPath", defaultConfigPath)

	rootCmd.PersistentFlags().StringVarP(&configFileName, "configFileName", "c", defaultConfigFileName, "Configuration file name.")
	rootCmd.PersistentFlags().StringVarP(&configPath, "configPath", "p", defaultConfigPath, "Configuration path.")
	v.BindPFlag("configFileName", rootCmd.PersistentFlags().Lookup("configFileName"))
	v.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("configPath"))

	tokenCmd.Flags().StringVar(&username, "username", "", "WD My Cloud Home user name.")
	tokenCmd.Flags().StringVar(&password, "password", "", "WD My Cloud Home user password")
	tokenCmd.Flags().StringVar(&clientId, "clientId", "", "Client Id")
	tokenCmd.Flags().StringVar(&clientSecret, "clientSecret", "", "Client Secret")
	v.BindPFlag("username", tokenCmd.Flags().Lookup("username"))
	v.BindPFlag("password", tokenCmd.Flags().Lookup("password"))
	v.BindPFlag("clientId", tokenCmd.Flags().Lookup("clientId"))
	v.BindPFlag("clientSecret", tokenCmd.Flags().Lookup("clientSecret"))

	if cp := v.GetString("configPath"); cp != "" {
		if cf := v.GetString("configFileName"); cf != "" {
			addJsonFile(cp, cf)
		}
	}

	v.SetEnvPrefix("mch")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	cobra.OnInitialize(func() {
		addJsonFile(configPath, configFileName)
	})

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.Execute()
}

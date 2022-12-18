package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/utils"
	"github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

var (
	//home       = os.Getenv("HOME")
	absPath, _ = filepath.Abs(".")
)

const (
	defaultConfigFileName string = "config"
)

var (
	defaultConfigPath string = absPath
)

var rootCmd = &cobra.Command{
	Use:   "my-cloud-home-go",
	Short: "My Cloud Home CLI application",
	Long:  `my-cloud-home-go is meant for managing My Cloud Home devices`,
}

func InitRootCommand(v *viper.Viper) *cobra.Command {

	v.SetDefault("configFileName", defaultConfigFileName)
	v.SetDefault("configPath", defaultConfigPath)

	rootCmd.PersistentFlags().StringP("configPath", "p", defaultConfigPath, "Configuration path.")
	rootCmd.PersistentFlags().StringP("configFileName", "f", defaultConfigFileName, "Configuration file name.")
	rootCmd.MarkFlagsRequiredTogether("configPath", "configFileName")
	v.BindPFlag("configPath", rootCmd.PersistentFlags().Lookup("configPath"))
	v.BindPFlag("configFileName", rootCmd.PersistentFlags().Lookup("configFileName"))

	initConfig(v)

	return rootCmd
}

func initConfig(v *viper.Viper) {
	addJsonFile(v)
	cobra.OnInitialize(func() {
		if v.GetString("configPath") != defaultConfigPath && v.GetString("configFileName") != defaultConfigFileName {
			addJsonFile(v)
		}
	})
}

func addJsonFileByPath(v *viper.Viper, filePath string, fileName string) {
	const ext = "json"
	in := filePath + string(os.PathSeparator) + fileName + "." + ext
	if in != "" {
		if utils.FileExists(in) {
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

func addJsonFile(v *viper.Viper) {
	if cp := v.GetString("configPath"); cp != "" {
		if cf := v.GetString("configFileName"); cf != "" {
			addJsonFileByPath(v, cp, cf)
		}
	}
}

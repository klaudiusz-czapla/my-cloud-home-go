package main

import (
	"log"
	"path/filepath"

	"github.com/klaudiusz-czapla/my-cloud-home-go/cmd"
	"github.com/spf13/viper"
)

var (
	absPath, _ = filepath.Abs(".")
	v          = viper.New()
)

func main() {
	log.Print("app has been started..")
	log.Printf("started from the path: %s", absPath)

	rootCmd := cmd.InitRootCommand(v)
	versionCmd := cmd.InitVersionCommand()
	configCmd := cmd.InitConfigCommand()
	tokenCmd := cmd.InitTokenCommand(v)
	refreshTokenCmd := cmd.InitRefreshTokenCommand(v)
	jwtCmd := cmd.InitJwtCommand(v)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(tokenCmd)
	rootCmd.AddCommand(refreshTokenCmd)
	rootCmd.AddCommand(jwtCmd)

	tokenCmd.AddCommand(jwtCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}

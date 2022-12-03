package main

import (
	"fmt"
	"log"

	mch "github.com/klaudiusz-czapla/my-cloud-home-go/mch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {

	command := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}

	log.Println("Starting My Cloud Home client app..")

	viper.SetConfigType(".ini")
	viper.SetConfigName("config")

	clientId := viper.Get("clientid").(string)
	clientSecret := viper.Get("clientsecret").(string)

	log.Println("Configuration has been loaded..")

	var config *mch.MchConfig
	config, err := mch.GetConfiguration()
	log.Println(config.Data.ComponentMapUntyped)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Configuration retrieved with success..")

}

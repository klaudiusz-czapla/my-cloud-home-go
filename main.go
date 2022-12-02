package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	mch "github.com/klaudiusz-czapla/my-cloud-home-go/mch"
)

func main() {
	log.Println("Starting My Cloud Home client app..")

	resp, err := http.Get("https://config.mycloud.com/config/v1/config")
	if err != nil {
		log.Fatalln(err.Error())
	}

	respBytesArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var config mch.Config
	err = json.Unmarshal(respBytesArr, &config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Configuration retrieved with success..")

}

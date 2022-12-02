package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting My Cloud Home client app..")

	resp, err := http.Get("https://config.mycloud.com/config/v1/config")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	respBytesArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respStr = string(respBytesArr)
	err := json.Unmarshal(respBytesArr, data)
	if err != nil {
		fmt.Println(err.Error())
		//json: Unmarshal(non-pointer main.Request)
	}

	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	}
}

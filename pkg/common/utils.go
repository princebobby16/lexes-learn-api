package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server string
}

var AppConfig configuration

func initConfig()  {
	loadAppConfig()
}

func loadAppConfig()  {
	file, err := os.Open("pkg/common/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s", err.Error())
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s", err.Error())
	}
}

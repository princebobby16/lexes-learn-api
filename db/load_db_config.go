package db

import (
	"encoding/json"
	"log"
	"os"
)

type pgInfo struct{
	Host string
	Port int
	User string
	Password string
	DbName string
}

var postgresConfig pgInfo

func LoadDbConfig() error {
	file, err := os.Open("data/db/db.json")
	if err != nil {
		log.Fatalf("[loadDbConfig]: %s", err.Error())
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&postgresConfig)
	if err != nil {
		log.Fatalf("[loadPgConfig]: %s", err.Error())
		return err
	}

	return nil
}

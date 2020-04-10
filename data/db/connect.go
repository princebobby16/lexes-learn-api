package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var err error

func Connect() {
	err = LoadDbConfig()
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("Unable to connect to database")
		panic(err)
	}

	log.Println("Connected to Postgres DB successfully")
}

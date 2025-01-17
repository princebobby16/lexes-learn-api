package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var err error

var DBConnection *sql.DB

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

	DBConnection = db

	err = db.Ping()
	if err != nil {
		log.Printf("Unable to connect to database")
		panic(err)
	}

	log.Println("Connected to Postgres DB successfully")
}

func DisconnectDB() error {
	err = DBConnection.Close()
	if err != nil {
		return err
	}

	return nil
}
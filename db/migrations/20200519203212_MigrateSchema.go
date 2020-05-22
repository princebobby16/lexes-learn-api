
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20200519203212(txn *sql.Tx) {
	_, err := txn.Exec("CREATE SCHEMA IF NOT EXISTS lexes AUTHORIZATION lexes;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = txn.Exec("CREATE SCHEMA IF NOT EXISTS ginart_academy AUTHORIZATION lexes;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Down is executed when this migration is rolled back
func Down_20200519203212(txn *sql.Tx) {
	_, err := txn.Exec("DROP SCHEMA IF EXISTS lexes, ginart_academy CASCADE;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

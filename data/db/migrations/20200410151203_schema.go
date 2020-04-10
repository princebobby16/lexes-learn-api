
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20200410151203(txn *sql.Tx) {
	_, err := txn.Exec("CREATE SCHEMA IF NOT EXISTS lexes AUTHORIZATION postgres;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Down is executed when this migration is rolled back
func Down_20200410151203(txn *sql.Tx) {
	_, err := txn.Exec("DROP SCHEMA IF EXISTS lexes CASCADE;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

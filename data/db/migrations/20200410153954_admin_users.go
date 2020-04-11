
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20200410153954(txn *sql.Tx) {
	_, err := txn.Exec("CREATE TABLE IF NOT EXISTS lexes.admin_users(" +
		"admin_id SERIAL PRIMARY KEY," +
		"institution_name VARCHAR(250)," +
		"institution_username VARCHAR(250)," +
		"institution_password VARCHAR(250)," +
		"digital_address VARCHAR(250)," +
		"mobile_number VARCHAR(250)," +
		"district VARCHAR(250)," +
		"time TIMESTAMP," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP) ;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Down is executed when this migration is rolled back
func Down_20200410153954(txn *sql.Tx) {
	_, err := txn.Exec("DROP TABLE IF EXISTS lexes.admin_users CASCADE ;")
	if err != nil {
		log.Fatal(err)
	}
	return
}

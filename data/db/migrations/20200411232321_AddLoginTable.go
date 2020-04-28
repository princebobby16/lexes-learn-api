
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20200411232321(txn *sql.Tx) {

	_, err := txn.Exec(`CREATE TABLE IF NOT EXISTS lexes.student
		(
			student_id character varying(100) NOT NULL UNIQUE,
    		lastname character varying(200) NOT NULL,
    		firstname character varying(200) NOT NULL,
    		username character varying(200) NOT NULL,
    		password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (student_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE lexes.student
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS lexes.teacher
		(
			teacher_id character varying(100) NOT NULL UNIQUE,
    		lastname character varying(200) NOT NULL,
    		firstname character varying(200) NOT NULL,
    		username character varying(200) NOT NULL,
    		password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (teacher_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE lexes.teacher
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}



	return
}

// Down is executed when this migration is rolled back
func Down_20200411232321(txn *sql.Tx) {
	_, err := txn.Exec(`
		DROP TABLE IF EXISTS lexes.student, lexes.teacher`)

	if err != nil {
		log.Println(err)
		return
	}
}

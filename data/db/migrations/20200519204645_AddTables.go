
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20200519204645(txn *sql.Tx) {

	_, err := txn.Exec(`CREATE TABLE IF NOT EXISTS lexes.ginart_academy
		(
			institution_id character varying(100) NOT NULL UNIQUE,
    		institution_name character varying(200) NULL,
    		institution_location character varying(200) NOT NULL,
    		institution_type character varying(200) NOT NULL,
    		institution_email character varying(200) NOT NULL UNIQUE,
    		institution_phone_number character varying(200) NOT NULL UNIQUE,
    		institution_alt_number character varying(200) NULL,
    		admin_name character varying(200) NOT NULL,
    		admin_username character varying(200) NOT NULL UNIQUE,
    		admin_password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (institution_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE lexes.ginart_academy
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.class
		(
			class_id character varying(100) NOT NULL UNIQUE,
    		class_name character varying(200) NOT NULL UNIQUE,
    		level character varying(200) NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (class_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.class
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.course
		(
			course_id character varying(100) NOT NULL UNIQUE,
    		course_name character varying(200) NOT NULL UNIQUE,
    		course_description character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (course_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.course
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.student
		(
			student_id character varying(100) NOT NULL UNIQUE,
    		lastname character varying(200) NOT NULL,
    		firstname character varying(200) NOT NULL,
    		date_of_birth character varying(200) NOT NULL,
    		phone_number character varying(200) NOT NULL UNIQUE,
    		class_enrolled_in character varying(200) NOT NULL REFERENCES gnart_academy.class,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (student_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.student
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.student_login
		(
			login_id character varying(100) NOT NULL REFERENCES ginart_academy.student,
    		username character varying(200) NOT NULL UNIQUE,
    		password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (login_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.student_login
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.teacher
		(
			teacher_id character varying(100) NOT NULL UNIQUE,
    		lastname character varying(200) NOT NULL,
    		firstname character varying(200) NOT NULL,
    		date_of_birth character varying(200) NOT NULL,
    		phone_number character varying(200) NOT NULL UNIQUE,
    		course character varying(200) REFERENCES ginart_academy.course,
    		class character varying(200) REFERENCES ginart_academy.class,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (teacher_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.teacher
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.teacher_login
		(
			login_id character varying(100) NOT NULL REFERENCES ginart_academy.teacher,
    		username character varying(200) NOT NULL UNIQUE,
    		password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (login_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.teacher_login
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.announcements
		(
			announcement_id character varying(100) NOT NULL UNIQUE,
    		title character varying(200) NULL,
    		content character varying(200) NOT NULL,
    		due_date timestamp with time zone NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (announcement_id)
		)
		WITH (
    		OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.announcements
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.multiple_choice_questions
		(
			question_id character varying (100) NOT NULL UNIQUE,
			course character varying (100) NOT NULL REFERENCES ginart_academy.course,
			question_type character varying (200) NOT NULL,
			question text NOT NULL,
			a_col character varying (200) NOT NULL,
			b_col character varying (200) NOT NULL,
			c_col character varying (200) NOT NULL,
			d_col character varying (200) NOT NULL, 
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (question_id)
		)
		WITH (
			OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.multiple_choice_questions
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.single_questions
		(
			question_id character varying (100) NOT NULL UNIQUE,
			course character varying (100) NOT NULL REFERENCES ginart_academy.course,
			question_type character varying (200) NOT NULL,
			question text NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (question_id)
		)
		WITH (
			OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.single_questions
    	OWNER to lexes;
		`)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = txn.Exec(`CREATE TABLE IF NOT EXISTS ginart_academy.assignments
		(
			assignment_id character varying(100) NOT NULL UNIQUE REFERENCES ginart_academy.class,
    		subject character varying(200),
    		course character varying(200) NOT NULL REFERENCES ginart_academy.course,
    		school character varying(200) NOT NULL REFERENCES lexes.ginart_academy,
    		duration time NOT NULL,
    		due_date time with time zone,
    		questions json NOT NULL,
    		score character varying(200),
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (assignment_id)
		)
		WITH (
			OIDS = FALSE
		);
	
		ALTER TABLE ginart_academy.assignments
    	OWNER to lexes;
		`)

	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}

// Down is executed when this migration is rolled back
func Down_20200519204645(txn *sql.Tx) {
	_, err := txn.Exec(`
		DROP TABLE IF EXISTS 
			ginart_academy.student, ginart_academy.teacher, 
			ginart_academy.announcements, ginart_academy.class,
			ginart_academy.teacher_login, ginart_academy.student_login,
			ginart_academy.multiple_choice_questions, ginart_academy.single_questions,
		    ginart_academy.assignments;`)

	if err != nil {
		log.Fatalln(err)
		return
	}
}

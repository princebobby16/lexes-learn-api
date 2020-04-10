package common

import "lexes_learn_server/data/db"

func StartUp() {
	initConfig()
	db.Connect()
}

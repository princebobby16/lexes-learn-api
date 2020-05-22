
package common

import "lexes_learn_server/db"

func StartUp() {
	initConfig()
	db.Connect()
}

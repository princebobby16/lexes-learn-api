package common

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)


func ComparePasswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


package common

import (
	"io/ioutil"
	"log"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

var err error

var (
	SignPrivateKey []byte
)

func Init() ([]byte, error) {
	SignPrivateKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return SignPrivateKey, nil
}

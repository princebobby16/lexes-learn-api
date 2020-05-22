package common

import (
	"io/ioutil"
	"log"
)

const (
	privateKeyPath = "pkg/keys/app.rsa"
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

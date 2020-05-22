package common

import (
	"github.com/dgrijalva/jwt-go"
	"log"
)

func DecodeToken(token string, signPrivateKey []byte) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return signPrivateKey, nil
	})
	if err != nil {
		log.Println("error")
		log.Println(err)
		return "", err
	}

	err = claims.Valid()
	if err != nil {
		return "", err
	}

	username, yes := claims["username"].(string)
	if !yes {
		log.Fatalln("Unable to get username")
	}

	return username, nil
}

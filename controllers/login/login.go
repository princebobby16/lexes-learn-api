package login

import (
	"fmt"
	"io/ioutil"
	"lexes_learn_server/models"
	"log"
	"net/http"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath = "keys/app.rsa.pub"
)

var err error

var (
	verifyPublicKey, signPrivateKey []byte
)

func init() {
	signPrivateKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	verifyPublicKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request)  {
	var loginRequest models.LoginRequest
	err := loginRequest.FromJson(r)
	if err != nil {
		// TODO: Handle error properly (Send appropriate error message)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}
	//	TODO: Validate user credentials
	//	TODO: Store user login detail log in the database
	//	TODO: Generate a token
	//	TODO: Send token to client
}

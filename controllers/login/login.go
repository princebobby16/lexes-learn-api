package login

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"lexes_learn_server/common"
	"lexes_learn_server/data/db"
	"lexes_learn_server/models"
	"log"
	"net/http"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
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

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	var userLogin models.LoginRequest

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			models.LoginErrorResponse{
				Status: "Error",
				Data: models.LoginErrorData{
					Code:    0,
					Message: "could not read request body",
				},
			},
		)
		return
	}

	log.Println(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &userLogin)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			models.LoginErrorResponse{
				Status: "",
				Data: models.LoginErrorData{
					Code:    0,
					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	// TODO: Check for empty or corrupt fields
	// check required fields
	/*err = validator.New().Struct(userLogin)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,

					Message: "bad data",
				},
			},
		)
		return
	}*/

	var (
		userData models.LoginData

		getUserData = `SELECT login_id, student_id, password
		FROM lexes.studentlogin
		WHERE username = $1`
	)

	// check username
	row, err := db.DBConnection.Query(getUserData, userLogin.Username)
	if err != nil {
		log.Println(err)
		return
	}

	err = row.Scan(
		&userData.LoginID,
		&userData.StudentID,
		&userData.Password,
	)
	if err != nil {
		log.Println(err)
		return
	}

	// compare passwords
	err = common.ComparePasswords(userLogin.Password, userData.Password)
	if err != nil {
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login_id":   userData.LoginID,
		"student_id": userData.StudentID,
		"password":   userData.Password,
	})

	tokenString, err := token.SignedString([]byte("way"))

	if err != nil {
		log.Println(err)
		return
	}

	// You can add expiration time to the token so that the person is not perpetually logged in
	// We may also need a logout function
	successResponse := models.LoginResponse{
		Token: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(successResponse)
	return

}

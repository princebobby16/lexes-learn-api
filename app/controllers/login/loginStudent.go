package login

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"lexes_learn_server/pkg/common"
	"lexes_learn_server/db"
	"lexes_learn_server/pkg/models"
	"log"
	"net/http"
)

func SignInStudentHandler(w http.ResponseWriter, r *http.Request) {

	signPrivateKey, err := common.Init()
	if err != nil {
		errorMessage := models.ErrorMessage{
			Status: "Failed",
			Message: "Internal Server Error",
		}
		log.Println(err)
		err = json.NewEncoder(w).Encode(errorMessage)
		if err != nil {
			log.Printf("[ErrorMessage]: %s", err.Error())
			return
		}
		return
	}

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

	//log.Println("Student")

	//log.Println(string(requestBody))

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

	namespace := common.GetSchemaName(userLogin.Username)
	log.Println(namespace)

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

		getUserData = fmt.Sprintf("SELECT login_id, username, password FROM %s.student_login WHERE username = $1", namespace)
	)

	// check username
	row, err := db.DBConnection.Query(getUserData, userLogin.Username)
	if err != nil {
		log.Println(err)
		return
	}

	ok := row.Next()
	if !ok {
		log.Println("Issue")
	}

	err = row.Scan(
		&userData.StudentID,
		&userData.Username,
		&userData.Password,
	)
	if err != nil {
		//log.Println("DB error")
		log.Println(err)
		log.Println("DB error")
		return
	}

	// compare passwords
	err = common.ComparePasswords(userLogin.Password, userData.Password)
	if err != nil {
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": userData.StudentID,
		"username": userData.Username,
		"password":   userData.Password,
		//"exp": time.Now().Add(time.Minute * 500).Unix(),
	})

	tokenString, err := token.SignedString(signPrivateKey)

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

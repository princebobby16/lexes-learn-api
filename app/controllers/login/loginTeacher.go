package login

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"lexes_learn_server/pkg/common"
	"lexes_learn_server/db"
	"lexes_learn_server/pkg/models"
	"log"
	"net/http"
	"time"
)


func SignInTeacherHandler(w http.ResponseWriter, r *http.Request) {
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
	// Create an instance of a request Object
	var request models.LoginRequest
	// Decode json Object into the request Object
	err = request.FromJson(r)

	// Handle error
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

	log.Println("Teacher")

	log.Println(request)

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

	namespace := common.GetSchemaName(request.Username)
	log.Println(namespace)

	var (
		userData models.LoginTeacherData

		getUserData = fmt.Sprintf("SELECT login_id, username, password FROM %s.teacher_login WHERE username = $1", namespace)
	)

	// check username
	row, err := db.DBConnection.Query(getUserData, request.Username)
	if err != nil {
		log.Println(err)
		return
	}

	ok := row.Next()
	if !ok {
		log.Println("Issue")
	}

	err = row.Scan(
		&userData.TeacherID,
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
	err = common.ComparePasswords(request.Password, userData.Password)
	if err != nil {
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"teacher_id": userData.TeacherID,
		"username": userData.Username,
		"password":   userData.Password,
		"exp": time.Now().Add(time.Minute * 20).Unix(),
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

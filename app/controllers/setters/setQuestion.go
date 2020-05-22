package setters

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"lexes_learn_server/pkg/common"
	"lexes_learn_server/db"
	"lexes_learn_server/pkg/models"
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

func SetQuestion(w http.ResponseWriter, r *http.Request)  {
	var requestData *models.QuestionsRequest

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(requestBody))

	err = json.Unmarshal(requestBody, &requestData)
	if err != nil {
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	log.Println(requestData)

	username, err := common.DecodeToken(requestData.Token, signPrivateKey)
	if err != nil {
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	var query string

	namespace := common.GetSchemaName(username)
	log.Println(namespace)

	// Generate uuid

	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: "Unable to generate uuid for course",
		})
		return
	}
	log.Println(id)

	if requestData.Question.QuestionType == "multiple_choice" {
		query = fmt.Sprintf("INSERT INTO %s.multiple_choice_questions(question_id, question_type, course, question, a_col, b_col, c_col, d_col) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", namespace)

		log.Println(query)

		result, err := db.DBConnection.Exec(query,
			id,
			requestData.Question.QuestionType,
			requestData.Question.CourseId,
			requestData.Question.Question,
			requestData.Question.ACol,
			requestData.Question.BCol,
			requestData.Question.CCol,
			requestData.Question.DCol,
		)
		if err != nil {
			_= json.NewEncoder(w).Encode(struct {
				Code uint
				Message string
			}{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(rowsAffected)
	}else if requestData.Question.QuestionType == "single" {
		query = fmt.Sprintf("INSERT INTO %s.single_questions(question_id, question_type, course, question) VALUES($1, $2, $3, $4)", namespace)
		log.Println(query)

		result, err := db.DBConnection.Exec(query,
			id,
			requestData.Question.QuestionType,
			requestData.Question.CourseId,
			requestData.Question.Question,
		)
		if err != nil {
			_= json.NewEncoder(w).Encode(struct {
				Code uint
				Message string
			}{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(rowsAffected)
	}else {
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: "Unacceptable request Object",
		})
		return
	}

	err = json.NewEncoder(w).Encode(struct {
		Status string
		Data struct {
			Code uint
			Message string
		}
	}{
		Status: "SUCCESS",
		Data: struct {
			Code    uint
			Message string
		}{
			Code: http.StatusOK,
			Message: "Successfully Set Question",
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
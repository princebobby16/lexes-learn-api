package getters

import (
	"encoding/json"
	"fmt"
	"lexes_learn_server/common"
	"lexes_learn_server/data/db"
	"lexes_learn_server/models"
	"log"
	"net/http"
)

func GetAllAssignments(w http.ResponseWriter, r *http.Request){
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
	var token *models.LoginResponse

	err = json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		log.Fatalln(err)
		return
	}

	username, err := common.DecodeToken(token.Token, signPrivateKey)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}

	namespace := common.GetSchemaName(username)
	log.Println(namespace)

	query := fmt.Sprintf("SELECT assignment_id, subject, course, school, duration, due_date, score, questions FROM %s.assignments", namespace)
	log.Println(query)

	rows, err:= db.DBConnection.Query(query)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: rows.Err().Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	var ann models.Assignments

	ok := rows.Next()
	if !ok {
		log.Println("Issue")
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: "No data found in DB",
			Code: http.StatusInternalServerError,
		})
	}
	err = rows.Scan(
		&ann.Id,
		&ann.Subject,
		&ann.Course,
		&ann.School,
		&ann.Duration,
		&ann.DueDate,
		&ann.Score,
		&ann.Questions,
	)

	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	err = json.NewEncoder(w).Encode(&ann)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(struct {
			Message string
			Code uint
		}{
			Message: err.Error(),
			Code: http.StatusInternalServerError,
		})

		return
	}
}

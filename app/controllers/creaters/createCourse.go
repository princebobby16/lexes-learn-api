package creaters

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

// CreateCourse Handle create corse 
func CreateCourse(w http.ResponseWriter, r *http.Request) {

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

	// Get request object
	var courseRequest *models.CourseRequest
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: "Unable to read request",
		})
		return
	}
	log.Println(string(requestData))

	err = json.Unmarshal(requestData, &courseRequest)
	if err != nil {
		log.Println(err)
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: "Unacceptable request Object",
		})
		return
	}

	log.Println(courseRequest)

	username, err := common.DecodeToken(courseRequest.Token, signPrivateKey)
	if err != nil {
		log.Println(err)
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusInternalServerError,
			Message: "Token error",
		})
		return
	}
	log.Println(username)

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

	query := fmt.Sprintf("INSERT INTO %s.course (course_id, course_name, course_description) VALUES ($1, $2, $3)", namespace)
	log.Println(query)

	result, err := db.DBConnection.Exec(query, id, courseRequest.Course.CourseName, courseRequest.Course.CourseDescription)
	if err != nil {
		log.Println(err)
		_= json.NewEncoder(w).Encode(struct {
			Code uint
			Message string
		}{
			Code: http.StatusConflict,
			Message: "Course Already Exists",
		})
		return
	}
	rowsAffected, _:= result.RowsAffected()
	log.Println(rowsAffected)

	err = json.NewEncoder(w).Encode(struct {
		Status string 		`json:"status"`
		Code uint			`json:"code"`
		Message string 		`json:"message"`
		CourseId uuid.UUID 	`json:"course_id"`
	}{
		Status: "SUCCESS",
		Code: http.StatusOK,
		Message: "Course Created",
		CourseId: id,
	})
}

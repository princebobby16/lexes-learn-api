package getters

import (
	"encoding/json"
	"fmt"
	"lexes_learn_server/pkg/common"
	"lexes_learn_server/db"
	"lexes_learn_server/pkg/models"
	"log"
	"net/http"
)


func GetAnnouncements(w http.ResponseWriter, r *http.Request) {

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
	// Get request
	var token *models.LoginResponse

	// Decode request
	err = json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(token)

	// Decode token
	username, err := common.DecodeToken(token.Token, signPrivateKey)
	if err != nil {
		log.Println(err)
		return
	}


	log.Println(username)

	namespace := common.GetSchemaName(username)
	log.Println(namespace)

	query := fmt.Sprintf("SELECT announcement_id, title, content, due_date FROM %s.announcements LIMIT $1", namespace)

	rows, err:= db.DBConnection.Query(query, 3)
	if err != nil {
		log.Println(rows.Err())
		log.Println(err)
		return
	}
	defer rows.Close()

	var ann models.Announcement

	var a []models.Announcement

	for rows.Next() {
		err = rows.Scan(
			&ann.Id,
			&ann.Title,
			&ann.Content,
			&ann.DueDate,
		)

		if err != nil {
			log.Println(err)
			return
		}

		a = append(a, ann)
	}
	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Println(err)
	}
}

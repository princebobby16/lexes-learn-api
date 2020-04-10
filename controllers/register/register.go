package register

import (
	"encoding/json"
	"lexes_learn_server/models"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Create an instance of a request Object
	var request models.RegisterRequest
	// Decode json Object into the request Object
	err := request.FromJson(r)

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

	//  TODO: Check if user already exists
	//	TODO: Hash Password
	//	TODO: Store User details in Database
	//	TODO: Send token to client
}

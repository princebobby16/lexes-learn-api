package register

import (
	"encoding/json"
	"go/token"
	"lexes_learn_server/models"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterRequest
	err := request.FromJson(r)

	if err != nil {
		errorMessage := models.ErrorMessage{
			Type: "Register",
			Status: "Failed",
			StatusCode: http.StatusBadRequest,
			Message: "Internal Server Error",
		}

		log.Println(err)
		err = json.NewEncoder(w).Encode(errorMessage)

		if err != nil {
			log.Printf("[ErrorMessage]: %s", err.Error())
		}

		//w.WriteHeader(http.StatusInternalServerError)
	}

	response := &models.RegisterResponse{
		Status: "Success",
		JWT:    token.ASSIGN,
	}

	w.Header().Set("Content-Type", "application/json")
	err = response.ToJson(w)
	if err != nil {
		errorMessage := models.ErrorMessage{
			Type: "Register",
			Status: "Failed",
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(errorMessage)
		if err != nil {
			log.Printf("Unable to encode response [ErrorMessage]: %s", err.Error())
		}

		//w.WriteHeader(http.StatusInternalServerError)
	}
}

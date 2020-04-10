package models

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string 			`json:"username"`
	Password string 			`json:"password"`
}


func (req *LoginRequest) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return err
	}
	return nil
}

type LoginResponse struct {
	Token json.Token 			`json:"token"`
}

func (resp *LoginResponse) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	return nil
}

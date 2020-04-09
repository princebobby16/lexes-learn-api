package models

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	InstitutionName string				`json:"institution_name"`
	InstitutionLocation string 			`json:"institution_location"`
	InstitutionAddress string 			`json:"institution_address"`
	InstitutionUsername string			`json:"institution_username"`
	InstitutionPassword string 			`json:"institution_password"`
}

func (req *RegisterRequest) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return err
	}
	return nil
}

type RegisterResponse struct {
	Status string 			`json:"status"`
	JWT json.Token 			`json:"jwt"`
}

func (resp *RegisterResponse) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	return nil
}
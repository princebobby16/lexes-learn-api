package models

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginRequest) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return err
	}
	return nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (resp *LoginResponse) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	return nil
}

type LoginErrorResponse struct {
	Status string         `json:"status"`
	Data   LoginErrorData `json:"data"`
}

type LoginErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginData struct {
	StudentID string `json:"student_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type LoginTeacherData struct {
	TeacherID string 	`json:"teacher_id"`
	Username string 	`json:"username"`
	Password string 	`json:"password"`
}
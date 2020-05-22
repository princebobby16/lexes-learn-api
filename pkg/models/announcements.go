package models

import (
	"encoding/json"
	"net/http"
)

type Announcement struct {
	Id string 			`json:"id"`
	Title string 		`json:"title"`
	Content string 		`json:"content"`
	DueDate string 		`json:"due_date"`
}

func (a *Announcement) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Announcement) ToJson(w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		return err
	}

	return nil
}

type TokenContent struct {
	StudentID string
	Username string
	Password string
	ExpirationDate string
}
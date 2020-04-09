package models

type HealthCheck struct {
	ServerName string	`json:"server_name"`
	Author string		`json:"author"`
	Version string 		`json:"version"`
	Health string		`json:"health"`
}

type ErrorMessage struct {
	Type string 			`json:"type"`
	Status string 			`json:"status"`
	Message string 			`json:"message"`
	StatusCode int 			`json:"status_code"`
}
package models

type ClassDetails struct {
	ClassName string 			`json:"class_name"`
	ClassTeacher string 		`json:"class_teacher"`
}

type ClassRequest struct {
	Token string 			`json:"token"`
	Class ClassDetails 		`json:"class"`
}
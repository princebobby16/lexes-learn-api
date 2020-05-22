package models


type CourseDetail struct {
	CourseName string		`json:"course_name"`
	CourseDescription string `json:"course_description"`
}

type CourseRequest struct {
	Token string 			`json:"token"`
	Course CourseDetail 	`json:"course"`
}

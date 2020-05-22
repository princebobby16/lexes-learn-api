package models

type QuestionInfo struct {
	CourseId string 		`json:"course_id"`
	QuestionId int	 		`json:"question_id"`
	QuestionType string		`json:"question_type"`
	Question string 		`json:"question"`
	ACol string 			`json:"a_col"`
	BCol string 			`json:"b_col"`
	CCol string				`json:"c_col"`
	DCol string 			`json:"d_col"`
}

type QuestionsRequest struct {
	Token string			`json:"token"`
	Question QuestionInfo	`json:"question"`
}
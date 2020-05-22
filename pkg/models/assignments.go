package models

import (
	"time"
)

type Assignments struct {
	Id 					string
	School 				string
	Subject 			string
	Course 				string
	Duration 			time.Time
	DueDate 			time.Time
	Score 				string
	Questions 			[]string
}
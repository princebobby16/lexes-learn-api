package router

import (
	"lexes_learn_server/controllers/healthcheck"
	"lexes_learn_server/controllers/login"
	"net/http"
)

// Create a single route object
type Route struct {
	Name string
	Path string
	Method string
	Handler http.HandlerFunc
}

// Create an object of different routes
type Routes []Route

func InitRoutes() Routes {
	routes := Routes{
		Route{
			Name:    "Index",
			Path:    "/",
			Method:  http.MethodGet,
			Handler: healthcheck.IndexHandler,
		},
		Route{
			Name:    "LoginStudent",
			Path:    "/users/student/login",
			Method:  http.MethodPost,
			Handler: login.SignInStudentHandler,
		},
		Route{
			Name:    "LoginTeacher",
			Path:    "/users/teacher/login",
			Method:  http.MethodPost,
			Handler: login.SignInTeacherHandler,
		},
	}

	return routes
}
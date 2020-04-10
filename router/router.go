package router

import (
	"lexes_learn_server/controllers/healthcheck"
	"lexes_learn_server/controllers/register"
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
			Name:    "Login",
			Path:    "/users/login",
			Method:  http.MethodPost,
			Handler: nil,
		},
		Route{
			Name:    "Register",
			Path:    "/users/register",
			Method:  http.MethodPost,
			Handler: register.Register,
		},
	}

	return routes
}
package router

import "net/http"

// Create a single route object
type Route struct {
	Name string
	Path string
	Method string
	Handler http.Handler
}

// Create an object of different routes
type Routes struct {
	Routes Route
}

func (routes Routes) initRoutes() Routes {
	routes = Routes{
		Route{
			Name:    "",
			Path:    "",
			Method:  "",
			Handler: nil,
		},

	}
	return routes
}

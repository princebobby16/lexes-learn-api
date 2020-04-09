package router

import "net/http"

// Create a single route object
type Route struct {
	Name string
	Path string
	Method string
	Handler http.HandlerFunc
}

// Create an object of different routes
type Routes []Route

//func SetRoutes() *Routes {
//
//}

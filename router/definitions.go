package router

import (
	"handler"
	"net/http"
)

// Route represents a routing information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represents a collection of routes
type Routes []Route

var routes = Routes{
	Route{
		"index",
		"POST",
		"/",
		handler.ReadPost,
	},
	Route{
		"testing",
		"GET",
		"/test",
		handler.Test,
	},
}

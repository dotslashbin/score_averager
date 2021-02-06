package router

import (
	"handler"
	"net/http"
)

// Route represents as routing information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represents a collection of routes
type Routes []Route

// This allows flexibiltiy to add the posibility of other routes should they be needed in the future.
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

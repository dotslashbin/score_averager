package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Mux returns an instance of the gorilla mux router
func Mux() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// goes thorugh all the defined routes
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		//   handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

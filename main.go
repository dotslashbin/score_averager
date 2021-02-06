package main

import (
	"log"
	"net/http"
	"router"
)

// Entry point of the application
/**
 * This loads up the router and starts the webserver
 */
func main() {
	router := router.Mux()
	log.Fatal(http.ListenAndServe(":5000", router))
}

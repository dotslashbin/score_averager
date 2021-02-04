package main

import (
	"log"
	"net/http"
	"router"
)

// Entry point of the application
func main() {
	router := router.Mux()
	log.Fatal(http.ListenAndServe(":5000", router))
}

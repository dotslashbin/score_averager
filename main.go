package main

import (
	"handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.ReadPost)

	log.Fatal(http.ListenAndServe(":5000", router))
}

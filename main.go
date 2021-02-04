package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " 1Hello from your basic API. Now start work")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handler)

	http.ListenAndServe(":5000", router)
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Scores provides the strucutre of the JSON inpout
type Scores struct {
	Foo string
}

// ReadPost will recieve and rad the values from a post request
func ReadPost(writer http.ResponseWriter, request *http.Request) {
	var scores Scores
	err := json.NewDecoder(request.Body).Decode(&scores)

	if err != nil {
		fmt.Println("There was an error when parsing the input")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(scores.Foo)
}

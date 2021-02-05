package handler

import (
	"encoding/json"
	"helper"
	"net/http"
	"payload"
)

// ReadPost will recieve and rad the values from a post request
func ReadPost(writer http.ResponseWriter, request *http.Request) {

	var inputScores payload.InputScores
	err := json.NewDecoder(request.Body).Decode(&inputScores)

	if err != nil {
		error := []string{"Invalid JSON input"}
		helper.DisplayOutput(false, error, writer)
	}

	helper.ValidateInput(inputScores, writer)

	// controller := app.Controller{}
	// controller.Compute(inputScores)

}

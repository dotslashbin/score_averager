package handler

import (
	"app"
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

	hasValidInputs := helper.ValidateInput(inputScores, writer)

	if hasValidInputs == true {
		controller := app.Controller{}
		controller.Compute(inputScores)
	}
}

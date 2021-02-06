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

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&inputScores)

	if err != nil {
		error := []string{err.Error()}

		helper.DisplayOutput(false, error, writer)
	} else {
		hasValidInputs := helper.ValidateInput(inputScores, writer)

		if hasValidInputs == true {
			controller := app.Controller{}
			controller.Compute(inputScores)
		}
	}
}

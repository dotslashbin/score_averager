/**
 * This file contains methods that act as route handlers for the POST requests.
 *
 * I normally maintain one responsibility for each file to make it a little more readable.
 */
package handler

import (
	"app"
	"encoding/json"
	"fmt"
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

		// Triggers to send an error when it fails to satisfy the json decoder's requirements
		helper.DisplayOutput(false, nil, error, writer)
	} else {
		hasValidInputs := helper.ValidateInput(inputScores, writer)

		if hasValidInputs == true {
			controller := app.Controller{}
			summary, error := controller.Compute(inputScores)

			if error == nil {
				helper.DisplayOutput(true, summary, []string{}, writer)
			} else {

				errorString := fmt.Sprintf("%v", error)

				errorsToDisplay := []string{errorString}
				helper.DisplayOutput(false, nil, errorsToDisplay, writer)
			}
		}
	}
}

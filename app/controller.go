/**
 * This provides the controller layer of the app. It is responsibile for beign 
 * the middle man between the model and the one that generates the output. 
 * 
 * I strongly believe that controllers should only be middlemans, which means it should 
 * not be doing all the processing, but rather calling other resources to do it. 
 * 
 * I normally kepe my controller methods light.
 */
package app

import (
	"helper"
	"net/http"
	"payload"
)

// Controller acts as the median between all layers and the input
type Controller struct{}

// Compute fetches computation of averages
func (controller *Controller) Compute(inputScores payload.InputScores, writer http.ResponseWriter) {

	summaryOfAverages := getSummaries(inputScores)

	// Triggers to send the output out after processing
	helper.DisplayOutput(true, summaryOfAverages, []string{}, writer)
}

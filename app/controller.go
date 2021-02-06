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

	helper.DisplayOutput(true, summaryOfAverages, []string{}, writer)
}

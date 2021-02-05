package app

import (
	"fmt"
	"payload"
)

// Controller acts as the median between all layers and the input
type Controller struct{}

// Compute fetches computation of averages
func (controller *Controller) Compute(inputScores payload.InputScores) {

	summaryOfAverages := getSummaries(inputScores)

	fmt.Println("output this: " + summaryOfAverages)
}

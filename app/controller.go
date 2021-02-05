package app

import (
	"payload"
)

// Controller acts as the median between all layers and the input
type Controller struct{}

// Compute fetches computation of averages
func (controller *Controller) Compute(inputScores payload.InputScores) {
	// if len(inputScores.Scores.Others) == 0 {
	// 	fmt.Println("woeuroeu")
	// }
}

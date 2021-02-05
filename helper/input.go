package helper

import (
	"model"
	"payload"
)

// getInputMap builds and returns a map out of the payload input
func getInputMap(inputScores payload.InputScores) map[string][]model.MemberScore {
	inputMap := make(map[string][]model.MemberScore)
	inputMap["managers"] = inputScores.Scores.Managers
	inputMap["team"] = inputScores.Scores.Team
	inputMap["others"] = inputScores.Scores.Others

	return inputMap
}

package helper

import (
	"model"
	"payload"
)

// GetInputMap builds and returns a map out of the payload input
func GetInputMap(inputScores payload.InputScores) map[string][]model.MemberScore {
	inputMap := make(map[string][]model.MemberScore)
	inputMap["managers"] = inputScores.Scores.Managers
	inputMap["team"] = inputScores.Scores.Team
	inputMap["others"] = inputScores.Scores.Others

	return inputMap
}

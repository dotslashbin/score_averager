/**
 * This file contains methods that will help for the input. 
 * 
 * It may not seem like it because there is only one method here, but I would normally 
 * group functinos in a file, if they are for the same purpose or a general group if they 
 * cannot be classified. 
 */
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

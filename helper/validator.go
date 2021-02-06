package helper

import (
	"model"
	"net/http"
	"payload"
	"strconv"
)

// hasSufficientData checks to see if there is at least one group that has data for processing
func hasSufficientData(mappedInput map[string][]model.MemberScore, writer http.ResponseWriter) bool {

	hasData := false

	for _, value := range mappedInput {
		if len(value) > 0 {
			hasData = true
		}
	}

	if hasData == false {
		error := []string{"There are no sufficient data for processing"}
		DisplayOutput(false, nil, error, writer)

		return false
	}

	return true
}

// hasValidData Checks to see if the supplied scoring data are valid
func hasValidData(mappedInput map[string][]model.MemberScore, writer http.ResponseWriter) bool {

	evaluatedUserids := []int{}

	for groupName, value := range mappedInput {
		for _, memberScore := range value {

			/**
			 * This validates if the given score are within range
			 */
			if memberScore.Score <= 0 || memberScore.Score > 5 {
				error := []string{
					"Invalid score input",
					"User: " + strconv.Itoa(memberScore.Userid) + " of " + groupName + " has an invalid score of " + strconv.Itoa(memberScore.Score) + ".",
				}
				DisplayOutput(false, nil, error, writer)
				return false
			}

			/**
			 * This validates if the user id has already been used or recorded
			 */
			if intInSlice(memberScore.Userid, evaluatedUserids) == false {
				evaluatedUserids = append(evaluatedUserids, memberScore.Userid)
			} else {
				error := []string{
					"Duplicate User ID detected",
					"User: " + strconv.Itoa(memberScore.Userid) + " has recorded a score.",
				}
				DisplayOutput(false, nil, error, writer)
				return false
			}
		}
	}

	return true
}

// ValidateInput executes the process of validating the input payload
func ValidateInput(inputScores payload.InputScores, writer http.ResponseWriter) bool {
	inputMap := GetInputMap(inputScores)
	if hasSufficientData(inputMap, writer) == true {
		return hasValidData(inputMap, writer)
	}

	return true
}

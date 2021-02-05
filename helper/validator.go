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
		DisplayOutput(false, error, writer)

		return false
	}

	return true
}

// hasValidData Checks to see if the supplied scoring data are valid
func hasValidData(mappedInput map[string][]model.MemberScore, writer http.ResponseWriter) {

	for groupName, value := range mappedInput {
		for _, memberScore := range value {
			if memberScore.Score <= 0 || memberScore.Score > 5 {
				error := []string{
					"Invalid  score input",
					"User: " + strconv.Itoa(memberScore.Userid) + " of " + groupName + " has an invalid score of " + strconv.Itoa(memberScore.Score),
				}
				DisplayOutput(false, error, writer)
			}
		}
	}
}

// ValidateInput executes the process of validating the input payload
func ValidateInput(inputScores payload.InputScores, writer http.ResponseWriter) {

	inputMap := getInputMap(inputScores)

	if hasSufficientData(inputMap, writer) == true {
		hasValidData(inputMap, writer)
	}

}

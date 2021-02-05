package app

import (
	"fmt"
	"helper"
	"model"
	"payload"
)

// getSummaries returns the summaries of scores based on groups
func getSummaries(inputScores payload.InputScores) string {
	inputMap := helper.GetInputMap(inputScores)

	for key, value := range inputMap {

		if key == "managers" {
			groupScore := model.ManagersScore{}
			groupScore.Name = key
			groupScore.ScoreCollection = value
			fmt.Println(groupScore)
			fmt.Println(groupScore.GetAverageScore())
		} else {
			groupScore := model.NonManagersScore{}
			groupScore.Name = key
			groupScore.ScoreCollection = value
			fmt.Println(groupScore)
			fmt.Println(groupScore.GetAverageScore())
		}
	}

	return "END OF PROCES"
}

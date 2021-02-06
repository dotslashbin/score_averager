package app

import (
	"helper"
	"model"
	"payload"
)

// getSummaries returns the summaries of scores based on groups
func getSummaries(inputScores payload.InputScores) map[string]float32 {
	inputMap := helper.GetInputMap(inputScores)

	summaries := make(map[string]float32)

	for key, value := range inputMap {

		if key == "managers" {
			groupScore := model.ManagersScore{}
			groupScore.ScoreCollection = value
			summaries[key] = groupScore.GetAverageScore()

		} else {
			groupScore := model.NonManagersScore{}
			groupScore.ScoreCollection = value
			summaries[key] = groupScore.GetAverageScore()
		}
	}

	return summaries
}

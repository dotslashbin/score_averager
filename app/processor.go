/**
 * This file contains all the methods that are involved for the main purpose of the application.
 *
 * Each method is responsible for only one thing. and it is the controller that calls on to them
 * for execution.
 */
package app

import (
	"helper"
	"model"
	"payload"
	"strings"
)

// getSummaries returns the summaries of scores based on groups
func getSummaries(inputScores payload.InputScores) map[string]interface{} {
	inputMap := helper.GetInputMap(inputScores)

	computedScores := make(map[string]float32)

	for key, value := range inputMap {

		dataModel := new(model.DataModel)
		dataModel.Init(key, value)

		averager := dataModel.GetAverager()
		computedScores[strings.Title(key)] = averager.GetAverageScore()
	}

	summaries := wrapWithLabel("Scores", computedScores)

	return summaries
}

// wrapWithLabel wraps the data structure into a map that contains a label
func wrapWithLabel(label string, data map[string]float32) map[string]interface{} {
	wrapper := make(map[string]interface{})
	wrapper[label] = data

	return wrapper
}

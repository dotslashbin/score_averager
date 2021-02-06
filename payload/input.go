/**
 * This defines the allowable input payload that is expected to be received  by the API endpoint. 
 * 
 * I normally do it this way to control the structure of the data that gets in any API endpoints. 
 * This provides a layer of security and peace of mind as the API develops. 
 * 
 * Should there be a need to cater to a dynamically structured input, reading of the input will 
 * also be done in a differnt way.
 * 
 */
package payload

import "model"

// InputScores Represents the inputted scores
type InputScores struct {
	Scores ScoresCollection
}

// ScoresCollection represents teh collection of score groupingsa
type ScoresCollection struct {
	Managers []model.MemberScore
	Team     []model.MemberScore
	Others   []model.MemberScore
}

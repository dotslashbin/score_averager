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

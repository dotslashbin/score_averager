package model

// TeamScores represen a collection of scores
type TeamScores struct {
	Managers []UserScore
	Team     []UserScore
	Others   []UserScore
}

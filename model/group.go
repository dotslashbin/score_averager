package model

// GroupScores Represents the collection of group scores
type GroupScores struct {
	Manabers []MemberScore
	Team     []MemberScore
	Others   []MemberScore
}

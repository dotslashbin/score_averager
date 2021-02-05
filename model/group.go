package model

// GroupScores Represents the collection of group scores
// type GroupScores struct {
// 	Manabers []MemberScore
// 	Team     []MemberScore
// 	Others   []MemberScore
// }

// GroupScores Defines the grouping of of socres
type GroupScores struct {
	Grouping map[string][]MemberScore
}

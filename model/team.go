package model

// TeamScore represents a team score
type TeamScore struct {
	Name    string
	Members []MemberScore
}

// CollectiveTeamScore represents a collection of team scores
type CollectiveTeamScore struct {
	Scores []Team
}

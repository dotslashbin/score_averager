package model

// ManagersScore represents the group for managers
type ManagersScore struct {
	Name            string
	ScoreCollection []MemberScore
}

// NonManagersScore represents the non-managers group, which has some constraints
type NonManagersScore struct {
	Name            string
	ScoreCollection []MemberScore
}

// Averager is the interface that contains methods to get the average summaries
type Averager interface {
	GetAverageScores()
}

// GetAverageScore Returns the average score of a manager group
func (group *ManagersScore) GetAverageScore() float32 {
	return computerForAverage(group.ScoreCollection)
}

// GetAverageScore Returns the average score of a non-manager group
func (group *NonManagersScore) GetAverageScore() float32 {

	numberOfScores := len(group.ScoreCollection)

	// This is the special case for non-manager group
	if numberOfScores > 3 {
		return computerForAverage(group.ScoreCollection)
	}

	return 0
}

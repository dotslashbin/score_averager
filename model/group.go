/**
 * This file provides the structure of a group data model.
 *
 * IMPORTANT!!
 * For the sake of this exercise, I decided  to create two structures that represents
 * the different type of groups. I chose this approach to demonstrate the use and knowledge
 * of interfaces.
 *
 * It might be considered as over-engineering, and there are simpler ways to do it.
 */
package model

// ManagersScore represents the group for managers
type ManagersScore struct {
	ScoreCollection []MemberScore
}

// GetAverageScore Returns the average score of a manager group
func (group ManagersScore) GetAverageScore() float32 {
	return computeForAverage(group.ScoreCollection)
}

// NonManagersScore represents the non-managers group, which has some constraints
type NonManagersScore struct {
	ScoreCollection []MemberScore
}

// GetAverageScore Returns the average score of a non-manager group
func (group NonManagersScore) GetAverageScore() float32 {

	numberOfScores := len(group.ScoreCollection)

	minRequirement := getMinRecordsRequirement()

	// This is the special case for non-manager group
	if numberOfScores > minRequirement {
		return computeForAverage(group.ScoreCollection)
	}

	return 0
}

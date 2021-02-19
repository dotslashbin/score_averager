package model

// DataModel a structure that represents a group-score data
type DataModel struct {
	groupType string
	scores    []MemberScore
}

// Init initializes the structure
func (dataModel *DataModel) Init(key string, scores []MemberScore) {
	dataModel.groupType = key
	dataModel.scores = scores
}

// GetAverager returns the correct referrencing interface based on the given key
func (dataModel *DataModel) GetAverager() Averager {
	var averager Averager

	if dataModel.groupType == "managers" {
		averager = ManagersScore{
			ScoreCollection: dataModel.scores,
		}
	} else {
		averager = NonManagersScore{
			ScoreCollection: dataModel.scores,
		}
	}

	return averager
}

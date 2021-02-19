package model

// Averager is the interface that contains methods to get the average summaries
type Averager interface {
	GetAverageScore() float32
}

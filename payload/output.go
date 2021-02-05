package payload

// Output provides the structure of the JSON output
type Output struct {
	Success bool
	Data    *averages
	Errors  []string
}

// ScoreAverage defines the structure of final average output of scoress
type averages struct {
	Managers float32
	Team     float32
	Others   float32
}

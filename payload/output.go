package payload

// Output provides the structure of the JSON output
type Output struct {
	Success bool
	Data    *averages
	Errors  []string
}

// averages defines the structure of final average output of scoress
type averages struct {
	Scores AverageSummaries
}

type AverageSummaries struct {
	Managers string
	Team     string
	Others   string
}

package payload

// Output provides the structure of the JSON output
type Output struct {
	Success bool
	Data    interface{}
	Errors  []string
}

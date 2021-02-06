/**
 * This provides a structure of the payload that is returned from the API 
 * to the receiver. 
 * 
 * I would always opt to use a standardized structure for output to make it easier 
 * for front end development. 
 */
package payload

// Output provides the structure of the JSON output
type Output struct {
	Success bool
	Data    interface{}
	Errors  []string
}

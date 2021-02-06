/**
 * This is a file where I would put methods that help in processing anything. 
 * 
 * I do not plan to put methos here that are constructed to a specific structure. 
 */
package helper

// intInSlice checks to see if an integer is already in the collection
func intInSlice(needle int, list []int) bool {
	for _, value := range list {
		if value == needle {
			return true
		}
	}
	return false
}

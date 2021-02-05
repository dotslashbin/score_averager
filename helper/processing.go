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

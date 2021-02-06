/**
 * This is a file where I would create methods that are meant to interact wiht data models, and data models 
 * only. I would normally include that in the same pakcage. 
 * 
 * If this method was to compute an average of anything ( not restricted to the MemberScore data), 
 * it would reside on the helper package. 
 */
package model

// computeForAverage returns the average score based on the given collection of sores
func computeForAverage(collection []MemberScore) float32 {
	var total float32
	counter := 0

	// Collecting for total
	for _, data := range collection {
		counter++
		total += float32(data.Score)
	}

	// Computation for average
	average := total / float32(counter)
	return average
}

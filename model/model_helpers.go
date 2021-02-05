package model

// computerForAverage returns the average score based on the given collection of sores
func computerForAverage(collection []MemberScore) float32 {
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

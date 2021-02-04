package handler

import (
	"encoding/json"
	"net/http"
	"payload"
)

// Test is a test function
func Test(writer http.ResponseWriter, request *http.Request) {

	sampleErrors := []string{"one", "tow"}

	test := payload.Output{
		Success: false,
		Errors:  sampleErrors,
	}

	json.NewEncoder(writer).Encode(test)
}

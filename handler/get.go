package handler

import (
	"encoding/json"
	"net/http"
	"payload"
)

// Test is a test function
func Test(writer http.ResponseWriter, request *http.Request) {

	sampleErrors := []string{"one", "tow", "another"}

	test := payload.Output{
		Success: false,
		Errors:  sampleErrors,
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(writer).Encode(test)
}

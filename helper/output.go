package helper

import (
	"encoding/json"
	"net/http"
	"payload"
)

// DisplayOutput method that returns the output to the browser
func DisplayOutput(isSuccessful bool, data interface{}, errorMessages []string, writer http.ResponseWriter) {

	output := payload.Output{
		Success: isSuccessful,
		Data:    data,
		Errors:  errorMessages,
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(output)
}

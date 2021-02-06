/**
 * This file contains methods that will help for the output. 
 * 
 * It may not seem like it because there is only one method here, but I would normally 
 * group functinos in a file, if they are for the same purpose or a general group if they 
 * cannot be classified. 
 */
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

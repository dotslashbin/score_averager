package handler

import (
	"fmt"
	"net/http"
	"os"
)

// Test is a test function
func Test(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(os.Getenv("MAX_SCORE"))
}

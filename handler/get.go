/**
 * This file contains methods that act as route handlers for the GET requests. 
 * 
 * I normally maintain one responsibility for each file to make it a little more readable. 
 */
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

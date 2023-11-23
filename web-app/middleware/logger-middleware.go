package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.RequestURI(), time.Since(start)) }() // Important: Note: Defer keyword is used to delay the execution of a function or a statement until the nearby function returns. In simple words, defer will move the execution of the statement to the very end inside a function.

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

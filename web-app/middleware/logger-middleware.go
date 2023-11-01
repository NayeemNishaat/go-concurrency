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
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Handle unrecognized routes
			if r.URL.Path != "/" {
				http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

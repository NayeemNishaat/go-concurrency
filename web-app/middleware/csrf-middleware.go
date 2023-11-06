package middleware

import (
	"fmt"
	"net/http"
)

func Csrf() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			csrfToken := r.URL.Query().Get("csrfToken")

			if csrfToken == "" {
				csrfToken = r.Form.Get("csrfToken")
			}

			if token, err := r.Cookie("token"); err != nil || csrfToken != token.Value {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintln(w, "Forbidden")
				return
			}

			f(w, r)
		}
	}
}

package middleware

import (
	"net/http"
	"time"
)

func Csrf() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			csrfToken := r.URL.Query().Get("csrfToken")

			if csrfToken == "" {
				csrfToken = r.Form.Get("csrfToken")
			}

			if token, err := r.Cookie("token"); err != nil || csrfToken != token.Value {
				http.SetCookie(w, &http.Cookie{Name: "msg", Value: "Forbidden", Expires: time.Now().Add(time.Second)})
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}

			f(w, r)
		}
	}
}

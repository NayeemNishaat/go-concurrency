package controller

import (
	"net/http"
	"web/template"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	// ctx := context.WithValue(r.Context())

	template.Render(w, r, "welcome.page.gohtml", nil)
}

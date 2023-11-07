package controller

import (
	"net/http"
	"web/lib"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusInternalServerError)
	}

	lib.Render(w, r, "welcome.page.gohtml", &lib.TemplateData{CsrfToken: token.Value})
}

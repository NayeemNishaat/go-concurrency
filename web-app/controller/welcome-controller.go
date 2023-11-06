package controller

import (
	"net/http"
	"web/template"
)

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusInternalServerError)
	}

	template.Render(w, r, "welcome.page.gohtml", &template.TemplateData{CsrfToken: token.Value})
}

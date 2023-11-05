package controller

import (
	"net/http"
	"regexp"
	"web/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// Note: Handle unrecognized routes
	re := regexp.MustCompile("/js/*")
	if r.URL.Path != "/" && !re.Match([]byte(r.URL.Path)) {
		http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
		return
	}
	// fmt.Fprintln(w, "Something went wrong!")
	template.Render(w, r, "home.page.gohtml", nil)
}

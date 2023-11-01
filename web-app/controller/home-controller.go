package controller

import (
	"net/http"
	"web/templates"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// Handle unrecognized routes
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusPermanentRedirect)
		return
	}

	templates.Render(w, r, "home.page.gohtml", nil)
}

// token, err := lib.GenerateToken(0)

// if err != nil {
// 	fmt.Fprintln(w, "Something went wrong!")
// 	return
// }

// fmt.Fprintln(w, token)

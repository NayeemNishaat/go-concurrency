package controller

import (
	"net/http"
	"web/templates"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "home.page.gohtml", nil)
	// fmt.Fprintln(w, "hello world")
}

// token, err := lib.GenerateToken(0)

// if err != nil {
// 	fmt.Fprintln(w, "Something went wrong!")
// 	return
// }

// fmt.Fprintln(w, token)

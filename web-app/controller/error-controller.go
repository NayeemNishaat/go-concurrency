package controller

import (
	"net/http"
	"web/lib"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	msg, err := r.Cookie("errorMsg")

	if err != nil {
		if r.URL.Path == "/404" {
			lib.Render(w, r, "error.page.gohtml", &lib.TemplateData{Error: "This Route Is Unavailable!"})
		} else {
			lib.Render(w, r, "error.page.gohtml", &lib.TemplateData{Error: "Something Went Wrong!"})
		}
		return
	}

	lib.Render(w, r, "error.page.gohtml", &lib.TemplateData{Error: msg.Value})
}

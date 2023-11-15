package controller

import (
	"log"
	"net/http"
	"web/lib"
)

func (cfg *Config) PlanPage(w http.ResponseWriter, r *http.Request) {
	plans, err := cfg.Models.Plan.GetAll()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	token, err := r.Cookie("token")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusInternalServerError)
		return
	}

	lib.Render(w, r, "plans.page.gohtml", &lib.TemplateData{CsrfToken: token.Value, DataMap: dataMap})
}

func (cfg *Config) Subscribe(w http.ResponseWriter, r *http.Request) {}

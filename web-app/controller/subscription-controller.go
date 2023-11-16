package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
	"web/lib"
	"web/model"
)

func (cfg *Config) PlanPage(w http.ResponseWriter, r *http.Request) {
	plans, err := cfg.Models.Plan.GetAll()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	token, err := r.Cookie("token")
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
		return
	}

	lib.Render(w, r, "plans.page.gohtml", &lib.TemplateData{CsrfToken: token.Value, DataMap: dataMap})
}

func (cfg *Config) Subscribe(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	planId, err := strconv.Atoi(id)

	if err != nil {
		http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Plan Id not found!", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	plan, err := cfg.Models.Plan.GetOne(planId)

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Plan not found.")
		r = r.WithContext(ctx)

		lib.Render(w, r, "plans.page.gohtml", nil)
		return
	}

	user, ok := r.Context().Value(lib.User{}).(model.User)

	if !ok {
		http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Please log in first.", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

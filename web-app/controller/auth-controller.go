package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"web/lib"
	"web/template"
)

func RegisterMethodManager(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RegisterPage(w, r)
		return
	case http.MethodPost:
		Register(w, r)
		return
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	case http.MethodPatch:
		// Update an existing record partially.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func LoginMethodManager(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		LoginPage(w, r)
		// middleware.Chain(LoginPage, []middleware.Middleware{middleware.Logging()}).ServeHTTP(w, r)
		return
	case http.MethodPost:
		Cfg.Login(w, r)
		return
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	case http.MethodPatch:
		// Update an existing record partially.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	template.Render(w, r, "register.page.gohtml", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	//
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	template.Render(w, r, "login.page.gohtml", nil)
}

func (cfg *Config) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := cfg.Models.User.GetByEmail(email)

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		template.Render(w, r, "login.page.gohtml", nil)
		return
	}

	validPassword, err := user.PasswordMatches(password)

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		template.Render(w, r, "login.page.gohtml", nil)
		return
	}

	if !validPassword {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		template.Render(w, r, "login.page.gohtml", nil)
		return
	}

	token, err := lib.GenerateToken(uint(user.ID))

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		template.Render(w, r, "login.page.gohtml", nil)
		return
	}

	// Note: Preparing And Sending JSON
	/* json, err := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: token})

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		template.Render(w, r, "login.page.gohtml", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json) */

	http.Redirect(w, r, fmt.Sprintf("/welcome?token=%s", token), http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}

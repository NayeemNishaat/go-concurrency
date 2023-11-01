package controller

import (
	"net/http"
	"web/templates"
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
		Login(w, r)
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
	templates.Render(w, r, "register.page.gohtml", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	//
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "login.page.gohtml", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	//
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//
}

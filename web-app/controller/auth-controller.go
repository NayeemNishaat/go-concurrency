package controller

import (
	"net/http"
	"web/templates"
)

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

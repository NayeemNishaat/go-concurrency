package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"web/lib"
	"web/model"

	"github.com/golodash/galidator"
)

func RegisterMethodManager(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RegisterPage(w, r)
		return
	case http.MethodPost:
		Cfg.Register(w, r)
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
	lib.Render(w, r, "register.page.gohtml", nil)
}

type RegisterData struct {
	FirstName string `g:"required" required:"$field is required"`
	LastName  string `g:"required" required:"$field is required"`
	Password  string `g:"isStrong,required,max=50" isStrong:"$field should contain at least a special character, a number, a uppercase letter and minimum 8 characters long" required:"$field is required"`
	Email     string `g:"required,min=5" required:"$field is required"`
}

func (cfg *Config) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)

		ctx := context.WithValue(r.Context(), lib.Error{}, "Something Went Wrong")
		r = r.WithContext(ctx)

		lib.Render(w, r, "register.page.gohtml", nil)
		return
	}

	u := model.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    false,
		IsAdmin:   false,
	}

	customValidators := galidator.Validators{"isStrong": lib.ValidateStrongPass}
	validator := lib.GetCustomValidator(RegisterData{}, customValidators)
	error := validator.Validate(u)

	if error != nil {
		// fmt.Printf("%+v\n", error)
		// fmt.Printf("%#v\n", error)

		s, err := json.MarshalIndent(&error, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(s))

		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid data provided.")
		r = r.WithContext(ctx)

		lib.Render(w, r, "register.page.gohtml", nil)
		return
	}

	userId, err := u.Insert(u)

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Failed To Register")
		r = r.WithContext(ctx)

		lib.Render(w, r, "register.page.gohtml", nil)
		return
	}

	token, err := lib.GenerateToken(userId, true)

	if err != nil {
		log.Println(err)

		ctx := context.WithValue(r.Context(), lib.Error{}, "Something Went Wrong")
		r = r.WithContext(ctx)

		lib.Render(w, r, "register.page.gohtml", nil)
		return
	}

	msg := lib.Message{
		To:       []string{u.Email},
		Subject:  "Activate Your Accounr",
		Template: "confirmation-mail",
		Data:     map[string]any{"link": fmt.Sprintf("%s/activate?token=%s", os.Getenv("BASE_URL"), token)},
	}
	cfg.postMail(msg)

	ctx := context.WithValue(r.Context(), lib.Success{}, "Please check your email to activate your account!")
	r = r.WithContext(ctx)

	lib.Render(w, r, "home.page.gohtml", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	lib.Render(w, r, "login.page.gohtml", nil)
}

func (cfg *Config) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := cfg.Models.User.GetByEmail(email)

	msg := lib.Message{
		To:      []string{"nayeem@mail.com"},
		Subject: "Invalid Login Attempt",
		Data:    map[string]any{"message": "Someone tried to login with invalid creds."},
	}

	if err != nil {
		cfg.postMail(msg)

		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	validPassword, err := user.PasswordMatches(password)

	if err != nil {
		cfg.postMail(msg)

		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	if !validPassword {
		cfg.postMail(msg)

		ctx := context.WithValue(r.Context(), lib.Error{}, "Invalid Credentials")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	if !user.Active {
		ctx := context.WithValue(r.Context(), lib.Warning{}, "Please check your email to activate your account and try again!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	token, err := lib.GenerateToken(user.ID, false)

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	// Note: Preparing And Sending JSON
	/* json, err := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: token})

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json) */

	expires := time.Now().AddDate(0, 0, 1)
	secure, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	maxAge, err := strconv.ParseUint(os.Getenv("COOKIE_MAX_AGE"), 10, 32)

	if err != nil {
		fmt.Println(err)
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	ck := http.Cookie{
		Name:     os.Getenv("COOKIE_NAME"),
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		Path:     "/",
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Value:    token,
		MaxAge:   int(maxAge),
		Expires:  expires,
	}
	http.SetCookie(w, &ck)

	// http.Redirect(w, r, fmt.Sprintf("/welcome?token=%s", token), http.StatusSeeOther)
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	expires := time.Now()
	secure, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))

	if err != nil {
		ctx := context.WithValue(r.Context(), lib.Error{}, "Something went wrong!")
		r = r.WithContext(ctx)

		lib.Render(w, r, "login.page.gohtml", nil)
		return
	}

	ck := http.Cookie{
		Name:     os.Getenv("COOKIE_NAME"),
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		Path:     "/",
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Value:    "",
		MaxAge:   0,
		Expires:  expires,
	}
	http.SetCookie(w, &ck)
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}

func Activate(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(lib.UserId{}).(int)

	if !ok {
		http.SetCookie(w, &http.Cookie{Name: "msg", Value: "Invalid Token", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	activationToken, ok := r.Context().Value(lib.ActivationToken{}).(bool)
	if !ok || !activationToken {
		http.SetCookie(w, &http.Cookie{Name: "msg", Value: "Invalid Token", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	// http://localhost/activate?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI4MjAzODYsInVzZXJJZCI6MX0.QtjWCgl16uA-0v00auHy47ux9CR8aMKRf-kh7mFCAnU
	// http://localhost/activate?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmF0aW9uVG9rZW4iOnRydWUsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTcxMjgyNjAxMiwidXNlcklkIjowfQ.IuSu9mrRoGsLxV9KpOTkdwMJTg_AYVC0XDTfNWHXRL4
	u := &model.User{ID: userId}
	u, err := u.GetOne(u.ID)

	if err != nil {
		http.SetCookie(w, &http.Cookie{Name: "msg", Value: "User Not Found", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	fmt.Println(u)
}

package controller

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		log.Println(err)
		http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Plan Id not found!", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	plan, err := cfg.Models.Plan.GetOne(planId)

	if err != nil {
		log.Println(err)
		ctx := context.WithValue(r.Context(), lib.Error{}, "Plan not found.")
		r = r.WithContext(ctx)

		lib.Render(w, r, "plans.page.gohtml", nil)
		return
	}

	user, ok := r.Context().Value(lib.User{}).(*model.User)

	if !ok {
		log.Println("Failed to cast lib.User{} to *model.User")
		http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Please log in first.", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	cfg.Wg.Add(1)
	go func() {
		cfg.Wg.Done()

		invoice, err := cfg.GetInvoice(user, plan)
		if err != nil {
			cfg.ErrorChan <- err
			return
		}

		msg := lib.Message{
			To:       []string{user.Email},
			Subject:  "Invoice",
			Data:     map[string]any{"invoice": invoice},
			Template: "invoice",
		}

		cfg.PostMail(msg)
	}()

	cfg.Wg.Add(1)
	go func() {
		defer cfg.Wg.Done()

		pdf := cfg.GenerateManual(user, plan)

		err := pdf.OutputFileAndClose(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))

		if err != nil {
			cfg.ErrorChan <- err
			return
		}

		byteFile, err := os.ReadFile(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))

		if err != nil {
			cfg.ErrorChan <- err
			return
		}

		msg := lib.Message{
			To:      []string{user.Email},
			Subject: "Your Manual",
			Data:    map[string]any{"message": "Your manual is attached!"},
			Attachments: map[string][]byte{
				"Manual.pdf": byteFile,
			},
		}

		// cfg.ErrorChan <- errors.New("testing error chan") // Test:
		cfg.PostMail(msg)
	}()

	err = cfg.Models.Plan.SubscribeUserToPlan(*user, *plan)

	if err != nil {
		log.Println(err)
		http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Subscription request failed.", Expires: time.Now().Add(time.Second)})
		http.Redirect(w, r, "/plan", http.StatusSeeOther)
		return
	}

	u, err := cfg.Models.User.GetOne(user.ID)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	byteUser, err := json.Marshal(u)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	strUser := base64.StdEncoding.EncodeToString(byteUser)

	ck, err := lib.GenerateCookie("user", strUser, time.Time{})

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, ck)

	http.SetCookie(w, &http.Cookie{Name: "succMsg", Value: "Subscription Successful", Expires: time.Now().Add(time.Second)})
	http.Redirect(w, r, "/plan", http.StatusSeeOther)
}

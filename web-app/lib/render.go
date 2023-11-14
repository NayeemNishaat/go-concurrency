package lib

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"web/model"
)

var pathToTemplates = "./template"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	DataMap       map[string]any
	Success       string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	CsrfToken     string
	User          *model.User
}

func Render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplates),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))

	templateSlice = append(templateSlice, partials...)

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, AddDefaultData(td, r)); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	v, ok := r.Context().Value(Success{}).(string)
	if ok {
		td.Success = v
	} else {
		td.Success = ""
	}

	v, ok = r.Context().Value(Warning{}).(string)
	if ok {
		td.Warning = v
	} else {
		td.Warning = ""
	}

	v, ok = r.Context().Value(Error{}).(string)
	if ok {
		td.Error = v
	} else {
		td.Error = ""
	}

	_, ok = r.Context().Value(UserId{}).(uint64)

	if ok {
		td.Authenticated = true

		userCookie, err := r.Cookie("user")
		if err != nil {
			log.Println("User data not found.")
		} else {
			var user model.User
			err = json.Unmarshal([]byte(userCookie.Value), &user)

			if err != nil {
				log.Println("Failed to unmarshall data")
			} else {
				td.User = &user
			}
		}
	} else {
		td.Authenticated = false
	}

	td.Now = time.Now()

	return td
}

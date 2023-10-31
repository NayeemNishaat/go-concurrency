package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"web/middleware"
)

var pathToTemplates = "./templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	DataMap       map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	// User          *data.User
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
	td.Flash = "This Is Flash"
	td.Warning = "This Is Warning"
	td.Error = "This Is Error"

	_, ok := r.Context().Value(middleware.UserId{}).(int)
	if ok {
		td.Authenticated = true
	} else {
		td.Authenticated = false
	}

	td.Now = time.Now()

	return td
}

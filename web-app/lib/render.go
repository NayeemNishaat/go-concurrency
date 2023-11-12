package lib

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var pathToTemplates = "./template"

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
	CsrfToken     string
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
	td.Flash = ""
	td.Warning = ""

	v, ok := r.Context().Value(Success{}).(string)
	if ok {
		td.Flash = v
	} else {
		td.Flash = ""
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
	} else {
		td.Authenticated = false
	}

	td.Now = time.Now()

	return td
}

package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"web/model"
)

var PathToTemplates = "./template"

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
		fmt.Sprintf("%s/base.layout.gohtml", PathToTemplates),
		fmt.Sprintf("%s/header.partial.gohtml", PathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", PathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", PathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", PathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", PathToTemplates, t))

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
	msg, err := r.Cookie("succMsg")
	if err == nil {
		td.Success = msg.Value
	} else {
		v, ok := r.Context().Value(Success{}).(string)
		if ok {
			td.Success = v
		} /* else {
			td.Success = ""
		} */
	}

	msg, err = r.Cookie("warnMsg")
	if err == nil {
		td.Warning = msg.Value
	} else {
		v, ok := r.Context().Value(Warning{}).(string)
		if ok {
			td.Warning = v
		} /* else {
			td.Warning = ""
		} */
	}

	msg, err = r.Cookie("errorMsg")
	if err == nil {
		td.Error = msg.Value
	} else {
		v, ok := r.Context().Value(Error{}).(string)
		if ok {
			td.Error = v
		} /* else {
			td.Error = ""
		} */
	}

	_, ok := r.Context().Value(UserId{}).(int)

	if ok {
		td.Authenticated = true

		userCookie, err := r.Cookie("user")
		if err != nil {
			log.Println("User data not found.")
		} else {
			var user model.User

			decodedUser, err := base64.StdEncoding.DecodeString(userCookie.Value)

			if err != nil {
				log.Println("Failed to decode cookie")
			} else {
				err = json.Unmarshal(decodedUser, &user)

				if err != nil {
					log.Println("Failed to unmarshall data")
				} else {
					td.User = &user
				}
			}
		}
	} else {
		td.Authenticated = false
	}

	td.Now = time.Now()

	return td
}

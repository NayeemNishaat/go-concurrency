package lib

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	ctx := context.WithValue(req.Context(), Error{}, "Error")
	req = req.WithContext(ctx)

	req.AddCookie(&http.Cookie{Name: "succMsg", Value: "Success"})

	td := AddDefaultData(&TemplateData{}, req)

	if td.Error != "Error" {
		t.Error("Failed to get error data from ctx.")
	}

	if td.Success != "Success" {
		t.Error("Failed to get success data from cookie.")
	}
}

func TestIsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	_, ok := req.Context().Value(UserId{}).(int)

	if ok {
		t.Error("Should be unauthenticated.")
	}

	ctx := context.WithValue(req.Context(), UserId{}, 100)
	req = req.WithContext(ctx)

	_, ok = req.Context().Value(UserId{}).(int)

	if !ok {
		t.Error("Should be authentication")
	}
}

func TestRender(t *testing.T) {
	pathToTemplates = "../template"

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	Render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != http.StatusOK {
		t.Error("Failed to render home page")
	}

	// fmt.Println(rr.Body)
}

package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"
	"web/lib"
)

func TestLogin(t *testing.T) {
	lib.PathToTemplates = "../template"

	postedData := url.Values{
		"email":    {"23@mail.com"},
		"password": {"123456a"},
	}

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(postedData.Encode()))

	handler := http.HandlerFunc(TestConfig.Login)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("Wrong code returned!")
	}

	cookies := rr.Result().Cookies()

	tokenExist := slices.ContainsFunc(cookies, func(c *http.Cookie) bool {
		return c.Name == "token"
	})

	if !tokenExist {
		t.Error("Login Failed")
	}
}

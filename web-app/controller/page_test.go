package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web/lib"
)

func TestPages(t *testing.T) {
	lib.PathToTemplates = "../template"

	testCases := []struct {
		desc    string
		url     string
		status  int
		handler http.HandlerFunc
		html    string
	}{
		{
			desc:    "Home Page",
			url:     "/",
			status:  http.StatusOK,
			handler: HomePage,
		},
		{
			desc:    "Login Page",
			url:     "/login",
			status:  http.StatusOK,
			handler: LoginPage,
			html:    `<h1 class="mt-5">Login</h1>`,
		},
		{
			desc:    "Logout",
			url:     "/logout",
			status:  http.StatusPermanentRedirect,
			handler: Logout,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, tC.url, nil)

			handler := http.HandlerFunc(tC.handler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tC.status {
				t.Errorf("%s failed: Expected %d but got %d", tC.desc, tC.status, rr.Code)
			}

			if len(tC.html) > 0 {
				html := rr.Body.String()

				if !strings.Contains(html, tC.html) {
					t.Errorf("%s failed: Expected to find %s but didn't.", tC.desc, tC.html)
				}
			}
		})
	}
}

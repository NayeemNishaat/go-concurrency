package controller

import (
	"net/http"
	"net/http/httptest"
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
	}{
		{
			desc:    "Home Page",
			url:     "/",
			status:  http.StatusOK,
			handler: HomePage,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, tC.url, nil)

			handler := http.HandlerFunc(tC.handler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tC.status {
				t.Error("Expected 200 but got ", rr.Code)
			}
		})
	}
}

package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"web/lib"
	"web/model"
)

func TestSubscription(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/subscribe?id=1", nil)

	ctx := context.WithValue(req.Context(), lib.User{}, &model.User{ID: 1, Email: "admin@example.cc", FirstName: "Admin", LastName: "ACC", Active: true})
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(TestConfig.Subscribe)
	handler.ServeHTTP(rr, req)

	TestConfig.Wg.Wait()

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Expected %d but got %d", http.StatusSeeOther, rr.Code)
	}
}

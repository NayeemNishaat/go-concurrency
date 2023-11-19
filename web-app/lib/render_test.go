package lib

import (
	"context"
	"net/http"
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

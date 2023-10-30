package controller

import (
	"fmt"
	"net/http"
	"web/lib"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	token, err := lib.GenerateToken(0)

	if err != nil {
		fmt.Fprintln(w, "Something went wrong!")
		return
	}

	fmt.Fprintln(w, token)
}

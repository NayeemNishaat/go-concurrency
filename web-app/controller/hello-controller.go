package controller

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// token, err := lib.GenerateToken(0)

// if err != nil {
// 	fmt.Fprintln(w, "Something went wrong!")
// 	return
// }

// fmt.Fprintln(w, token)

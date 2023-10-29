package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)

	if r.URL.String() == "/" {
		w.Write([]byte("The Best Router!"))
		return
	}

	http.NotFound(w, r)
}

package middleware

import (
	"net/http"
	"slices"
)

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, globalMiddlewars []Middleware, middlewares ...Middleware) http.HandlerFunc {
	middlewares = append(globalMiddlewars, middlewares...)
	slices.Reverse[[]Middleware](middlewares)

	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

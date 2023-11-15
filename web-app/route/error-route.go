package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func ErrorRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/error", middleware.Chain(controller.ErrorPage, globalMiddlewars, middleware.Method("GET"), middleware.Token(true)))
	m.HandleFunc("/404", middleware.Chain(controller.ErrorPage, globalMiddlewars, middleware.Method("GET"), middleware.Token(true)))
}

package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func HelloRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/hello", middleware.Chain(controller.HomePage, globalMiddlewars, middleware.Method("GET"), middleware.Token()))
}

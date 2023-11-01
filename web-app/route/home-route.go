package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func HomeRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/", middleware.Chain(controller.HomePage, globalMiddlewars, middleware.Method("GET")))
}

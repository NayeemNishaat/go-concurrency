package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func WelcomeRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/welcome", middleware.Chain(controller.WelcomePage, globalMiddlewars, middleware.Method("GET"), middleware.Token()))
}

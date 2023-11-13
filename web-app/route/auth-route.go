package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func AuthRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/register", middleware.Chain(controller.RegisterMethodManager, globalMiddlewars))

	m.HandleFunc("/activate", middleware.Chain(controller.Activate, globalMiddlewars, middleware.Method("GET"), middleware.Token(false)))

	m.HandleFunc("/login", middleware.Chain(controller.LoginMethodManager, globalMiddlewars))

	m.HandleFunc("/logout", middleware.Chain(controller.Logout, globalMiddlewars, middleware.Method("GET"), middleware.Csrf(), middleware.Token(false)))
}

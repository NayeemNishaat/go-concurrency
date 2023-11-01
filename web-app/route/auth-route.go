package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func AuthRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/register", middleware.Chain(controller.RegisterPage, globalMiddlewars, middleware.Method("GET")))

	// m.HandleFunc("/register", middleware.Chain(controller.Register, globalMiddlewars, middleware.Method("POST")))

	m.HandleFunc("/login", middleware.Chain(controller.LoginPage, globalMiddlewars))

	// m.HandleFunc("/login", middleware.Chain(controller.Login, globalMiddlewars, middleware.Method("POST")))

	m.HandleFunc("/logout", middleware.Chain(controller.Logout, globalMiddlewars, middleware.Method("GET"), middleware.Token()))
}

// switch r.Method {
// case http.MethodGet:
//     // Serve the resource.
// case http.MethodPost:
//     // Create a new record.
// case http.MethodPut:
//     // Update an existing record.
// case http.MethodDelete:
//     // Remove the record.
// default:
//     http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// }

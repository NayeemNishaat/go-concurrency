package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func HelloRoute(m *http.ServeMux) {
	m.HandleFunc("/hello", middleware.Chain(controller.Hello, middleware.Method("GET"), middleware.Token()))
}

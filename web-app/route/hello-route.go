package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func HelloRoute(m *http.ServeMux) {
	m.HandleFunc("/", middleware.Chain(controller.Hello, middleware.Method("GET"), middleware.Logging()))
}

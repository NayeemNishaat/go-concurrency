package route

import (
	"net/http"
	"web/controller"
	"web/middleware"
)

func SubscriptionRoute(m *http.ServeMux, globalMiddlewars ...middleware.Middleware) {
	m.HandleFunc("/plan", middleware.Chain(controller.Cfg.PlanPage, globalMiddlewars, middleware.Method("GET"), middleware.Token(false)))

	m.HandleFunc("/subscribe", middleware.Chain(controller.Cfg.Subscribe, globalMiddlewars, middleware.Method("GET"), middleware.Csrf(), middleware.Token(false)))
}

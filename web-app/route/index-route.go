package route

import (
	"fmt"
	"net/http"
	"web/middleware"
)

func InitRouter(m *http.ServeMux) {
	HomeRoute(m, middleware.Logging())
	WelcomeRoute(m, middleware.Logging())
	AuthRoute(m, middleware.Logging())

	m.HandleFunc("/404", middleware.Chain(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "This Route Is Unavailable!\n") }, []middleware.Middleware{middleware.Logging()}))

	m.HandleFunc("/500", middleware.Chain(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Something Went Wrong!\n") }, []middleware.Middleware{middleware.Logging()}))
}

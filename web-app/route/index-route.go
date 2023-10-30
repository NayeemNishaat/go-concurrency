package route

import (
	"fmt"
	"net/http"
	"web/middleware"
)

func InitRouter(m *http.ServeMux) {
	HelloRoute(m, middleware.Logging())

	m.HandleFunc("/", middleware.Chain(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "This Route Is Unavailable!") }, []middleware.Middleware{middleware.Logging()}))
}

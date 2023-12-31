package route

import (
	"net/http"
	"web/middleware"
)

func InitRouter(m *http.ServeMux) {
	HomeRoute(m, middleware.Logging())
	WelcomeRoute(m, middleware.Logging())
	AuthRoute(m, middleware.Logging())
	SubscriptionRoute(m, middleware.Logging())
	ErrorRoute(m, middleware.Logging())

	/* m.HandleFunc("/error", middleware.Chain(func(w http.ResponseWriter, r *http.Request) {
		msg, err := r.Cookie("errorMsg")

		if err != nil {
			fmt.Fprint(w, "Something Went Wrong!")
			return
		}

		fmt.Fprint(w, msg.Value)
	}, []middleware.Middleware{middleware.Logging()})) */
}

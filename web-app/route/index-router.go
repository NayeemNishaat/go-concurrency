package route

import (
	"net/http"
)

func InitRouter(m *http.ServeMux) {
	HelloRoute(m)
}

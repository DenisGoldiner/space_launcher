package platform

import (
	"net/http"
)

func NewRouter(handlers map[string]http.Handler) http.Handler {
	r := http.NewServeMux()

	for route, handler := range handlers {
		r.Handle(route, handler)
	}

	return r
}

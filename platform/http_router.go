package platform

import (
	"fmt"
	"net/http"
)

const appBaseURL = "/space_launcher"

// NewRouter prepares the HTTP router.
func NewRouter(handlers map[string]http.Handler) http.Handler {
	r := http.NewServeMux()

	for route, handler := range handlers {
		r.Handle(BuildHandlerURL(route), handler)
	}

	return r
}

// BuildHandlerURL adds the appBaseURL prefix to handler url.
func BuildHandlerURL(route string) string {
	return fmt.Sprintf("%s%s", appBaseURL, route)
}

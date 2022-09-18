package platform

import (
	"fmt"
	"net/http"
)

const (
	appBaseURL = "/space_launcher"
)

func NewRouter(handlers map[string]http.Handler) http.Handler {
	r := http.NewServeMux()

	for route, handler := range handlers {
		r.Handle(buildHandlerURL(route), handler)
	}

	return r
}

func buildHandlerURL(route string) string {
	return fmt.Sprintf("%s%s", appBaseURL, route)
}

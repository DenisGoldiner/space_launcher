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
		r.Handle(BuildHandlerURL(route), handler)
	}

	return r
}

func BuildHandlerURL(route string) string {
	return fmt.Sprintf("%s%s", appBaseURL, route)
}

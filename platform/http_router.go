package platform

import (
	"github.com/DenisGoldiner/space_launcher/pkg/api"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type RouteExtender interface {
	Extend(r chi.Router) chi.Router
}

func NewRouter() http.Handler {
	r := chi.NewRouter()

	slh := api.SpaceLauncherHandler{}
	slh.Extend(r)

	return r
}

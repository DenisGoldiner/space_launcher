package integration

import (
	"github.com/DenisGoldiner/space_launcher/pkg/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func newTestRouter(t *testing.T) http.Handler {
	r := chi.NewRouter()

	slh := api.SpaceLauncherHandler{}
	slh.Extend(r)

	return r
}

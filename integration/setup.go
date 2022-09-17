package integration

import (
	"github.com/DenisGoldiner/space_launcher/internal/api"
	"net/http"
	"testing"
)

func newTestRouter(t *testing.T) http.Handler {
	slh := api.SpaceLauncherHTTPHandler{}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

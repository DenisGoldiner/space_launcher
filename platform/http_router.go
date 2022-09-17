package platform

import (
	"github.com/DenisGoldiner/space_launcher/internal/api"

	"net/http"
)

func NewRouter() http.Handler {
	slh := api.SpaceLauncherHTTPHandler{}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

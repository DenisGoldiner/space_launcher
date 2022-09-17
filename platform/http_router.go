package platform

import (
	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/service"

	"net/http"
)

func NewRouter() http.Handler {
	sls := service.SpaceLauncherService{}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

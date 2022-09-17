package platform

import (
	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/infra/repo"
	"github.com/DenisGoldiner/space_launcher/internal/service"

	"net/http"
)

func NewRouter() http.Handler {
	lr := repo.LauncherRepo{}
	sls := service.SpaceLauncherService{LaunchRepo: lr}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

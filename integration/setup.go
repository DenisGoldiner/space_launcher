package integration

import (
	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/infra/repo"
	"github.com/DenisGoldiner/space_launcher/internal/service"
	"net/http"
	"testing"
)

func newTestRouter(t *testing.T) http.Handler {
	ur := repo.UserRepo{}
	lr := repo.LauncherRepo{}
	sls := service.SpaceLauncherService{LaunchRepo: lr, UserRepo: ur}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

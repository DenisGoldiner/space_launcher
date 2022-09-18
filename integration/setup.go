package integration

import (
	"net/http"
	"testing"

	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/infra/adapter"
	"github.com/DenisGoldiner/space_launcher/internal/infra/repo"
	"github.com/DenisGoldiner/space_launcher/internal/service"
)

func newTestRouter(t *testing.T) http.Handler {
	ur := repo.UserRepo{}
	lr := repo.LaunchRepo{}
	sxc := adapter.SpaceXClient{}
	sls := service.SpaceLauncherService{LaunchRepo: lr, UserRepo: ur, SpaceXClient: sxc}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	r := http.NewServeMux()
	r.Handle("/bookings", slh)

	return r
}

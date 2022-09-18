package platform

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/infra/adapter"
	"github.com/DenisGoldiner/space_launcher/internal/infra/repo"
	"github.com/DenisGoldiner/space_launcher/internal/service"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

const runMigrationsPath = "./migrations"

// FailedToStartErr the HTTP server can't be started, preparation failed.
var FailedToStartErr = errors.New("failed to start the space_launcher")

// RunApp runs the space_launcher application and starts the HTTP server.
func RunApp() error {
	ctx, cancel := context.WithCancel(context.Background())

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		trapped := <-gracefulStop
		log.Printf("shutting down the space_launcher beacuse of %q signal\n", trapped.String())
		cancel()
	}()

	conf, err := LoadConfig()
	if err != nil {
		return pkg.WrapErr(err.Error(), FailedToStartErr)
	}

	dbCon, err := NewConnection(conf.DBConfig)
	if err != nil {
		return pkg.WrapErr(err.Error(), FailedToStartErr)
	}

	if err := MigrateUp(dbCon, runMigrationsPath); err != nil {
		return pkg.WrapErr(err.Error(), FailedToStartErr)
	}

	handlers := buildHandlers(dbCon)
	router := NewRouter(handlers)

	RunHTTPServer(ctx, router)

	return nil
}

func buildHandlers(dbCon *sqlx.DB) map[string]http.Handler {
	httpClient := pkg.NewDefaultClient()
	sxc := adapter.SpaceXClient{Client: httpClient}
	lr := repo.LaunchRepo{}
	ur := repo.UserRepo{}
	sls := service.SpaceLauncherService{DBCon: dbCon, LaunchRepo: lr, UserRepo: ur, SpaceXClient: sxc}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	return map[string]http.Handler{
		"/bookings": slh,
	}
}

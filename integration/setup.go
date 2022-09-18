package integration

import (
	"fmt"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"github.com/DenisGoldiner/space_launcher/platform"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"

	"github.com/DenisGoldiner/space_launcher/internal/api"
	"github.com/DenisGoldiner/space_launcher/internal/infra/adapter"
	"github.com/DenisGoldiner/space_launcher/internal/infra/repo"
	"github.com/DenisGoldiner/space_launcher/internal/service"
)

const (
	dsnEnvVar          = "DB_DSN"
	bookingsURL        = "/space_launcher/bookings"
	testMigrationsPath = "./../migrations"
)

func newTestRouter(dbCon *sqlx.DB) http.Handler {
	httpClient := pkg.NewDefaultClient()
	sxc := adapter.SpaceXClient{Client: httpClient}
	lr := repo.LaunchRepo{}
	ur := repo.UserRepo{}
	sls := service.SpaceLauncherService{DBCon: dbCon, LaunchRepo: lr, UserRepo: ur, SpaceXClient: sxc}
	slh := api.SpaceLauncherHTTPHandler{Service: sls}

	handlers := map[string]http.Handler{
		"/bookings": slh,
	}

	r := http.NewServeMux()

	for route, handler := range handlers {
		r.Handle(platform.BuildHandlerURL(route), handler)
	}

	return r
}

func setupDB(t testing.TB) *sqlx.DB {
	dsn := fmt.Sprintf("%s&intervalstyle=iso_8601&search_path=%s", getDSN(), t.Name())
	testDB, err := sqlx.Open(platform.DriverName, dsn)
	require.NoError(t, err)

	t.Cleanup(func() { _ = testDB.Close() })

	_, err = testDB.DB.Exec("DROP SCHEMA IF EXISTS " + t.Name() + " CASCADE")
	require.NoError(t, err)

	_, err = testDB.DB.Exec("CREATE SCHEMA " + t.Name())
	require.NoError(t, err)

	err = platform.MigrateUp(testDB, testMigrationsPath)
	require.NoError(t, err)

	return testDB
}

func getDSN() string {
	dsn := "postgres://sl_postgres:password@localhost:5432/space_launcher?sslmode=disable"
	envValue, exists := os.LookupEnv(dsnEnvVar)
	if exists {
		dsn = envValue
	}

	return dsn
}

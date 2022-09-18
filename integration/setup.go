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

func setupDB(t *testing.T) *sqlx.DB {
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

	setupFixtures(t, testDB)

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

func setupFixtures(t *testing.T, dbExec *sqlx.DB) {
	userFixtureQuery := `
		INSERT INTO "user" (first_name, last_name, gender, birthday)
		VALUES ('John', 'Smith', 'male', '1999-01-08') RETURNING id`

	var userID string

	err := dbExec.QueryRowx(userFixtureQuery).Scan(&userID)
	require.NoError(t, err)

	launchFixtureQuery := `
		INSERT INTO "launch" (launchpad_id, destination, launch_date, user_id)
		VALUES 
		    ('5e9e4501f509094ba4566f84', 'Mars', '2021-01-01', $1),
		    ('5e9e4502f509092b78566f87', 'Pluto', '2021-01-07', $1)`

	_, err = dbExec.Exec(launchFixtureQuery, userID)
	require.NoError(t, err)
}

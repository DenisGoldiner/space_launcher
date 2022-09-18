package platform

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

// MigrateDBErr describes an error when we failed to apply DB migrations.
var MigrateDBErr = errors.New("failed to run migrate up")

// MigrateUp will run the Up migrations
func MigrateUp(db *sqlx.DB, migrationPath string) error {
	m, err := setupMigrate(db, migrationPath)
	if err != nil {
		return pkg.WrapErr(err.Error(), MigrateDBErr)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}

		return pkg.WrapErr(err.Error(), MigrateDBErr)
	}

	return nil
}

func setupMigrate(db *sqlx.DB, migrationPath string) (*migrate.Migrate, error) {
	conn, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, pkg.WrapErr("failed to connect to the database", err)
	}

	source := fmt.Sprintf("file://%s", migrationPath)

	m, err := migrate.NewWithDatabaseInstance(source, DriverName, conn)
	if err != nil {
		return nil, pkg.WrapErr("failed to create a migrate instance", err)
	}

	return m, nil
}

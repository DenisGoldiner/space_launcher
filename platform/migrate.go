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

const migrationsPath = "./migrations"

// MigrateUp will run the migrations
func MigrateUp(db *sqlx.DB) error {
	m, err := setupMigrate(db)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return pkg.WrapErr("failed to run migrate up", err)
	}

	return nil
}

func setupMigrate(db *sqlx.DB) (*migrate.Migrate, error) {
	conn, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, pkg.WrapErr("failed to connect to the database", err)
	}

	source := fmt.Sprintf("file://%s", migrationsPath)
	m, err := migrate.NewWithDatabaseInstance(source, driverName, conn)
	if err != nil {
		return nil, pkg.WrapErr("failed to create a migrate instance", err)
	}
	return m, nil
}

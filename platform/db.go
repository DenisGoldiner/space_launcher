package platform

import (
	"errors"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"github.com/jmoiron/sqlx"
)

// DriverName sqlx driver to be used.
const DriverName = "postgres"

// DBConnectionErr describes an error when failed to setup the DB connection.
var DBConnectionErr = errors.New("failed to connect to the DB")

// NewConnection create a connection to a postgres
func NewConnection(conf DBConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Open(DriverName, conf.DSN)
	if err != nil {
		return nil, pkg.WrapErr(err.Error(), DBConnectionErr)
	}

	conn.SetConnMaxLifetime(conf.MaxConnLifetime)
	conn.SetMaxOpenConns(conf.MaxOpenConnection)
	conn.SetMaxIdleConns(conf.MaxIdleConnection)

	return conn, nil
}

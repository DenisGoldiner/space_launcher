package platform

import (
	"github.com/jmoiron/sqlx"
)

const driverName = "postgres"

// NewConnection create a connection to a postgres
func NewConnection(conf DBConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Open(driverName, conf.DSN)
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(conf.MaxConnLifetime)
	conn.SetMaxIdleConns(conf.MaxIdleConnection)
	conn.SetMaxOpenConns(conf.MaxOpenConnection)

	return conn, nil
}

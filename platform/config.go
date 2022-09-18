package platform

import (
	"errors"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

// ParseConfigErr describes an error when `envconfig` failed to process the config.
var ParseConfigErr = errors.New("failed to parse the config")

// Config keeps all the configurations together.
type Config struct {
	DBConfig
}

// DBConfig is used to setup the postgreSQL DB connection.
type DBConfig struct {
	DSN               string        `envconfig:"DB_DSN" required:"true"`
	MaxConnLifetime   time.Duration `envconfig:"MAX_CONNECTION_LIFETIME" default:"10s"`
	MaxOpenConnection int           `envconfig:"MAX_OPEN_CONNECTION" default:"10"`
	MaxIdleConnection int           `envconfig:"MAX_IDLE_CONNECTION" default:"10"`
}

// LoadConfig can be used to process configuration setup.
func LoadConfig() (Config, error) {
	var conf Config

	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, pkg.WrapErr(err.Error(), ParseConfigErr)
	}

	return conf, nil
}

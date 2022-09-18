package platform

import (
	"errors"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

var (
	ParseConfigErr = errors.New("failed to parse the config")
)

type Config struct {
	DBConfig
}

type DBConfig struct {
	DSN               string        `envconfig:"DB_DSN" required:"true"`
	MaxOpenConnection int           `envconfig:"MAX_OPEN_CONNECTION" default:"10"`
	MaxIdleConnection int           `envconfig:"MAX_IDLE_CONNECTION" default:"2"`
	MaxConnLifetime   time.Duration `envconfig:"MAX_CONNECTION_LIFETIME" default:"10s"`
}

// LoadConfig can be used to retrieve api configuration setup
func LoadConfig() (Config, error) {
	var conf Config

	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, pkg.WrapErr(err.Error(), ParseConfigErr)
	}

	return conf, nil
}

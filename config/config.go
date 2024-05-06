package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	ApplicationPort int  `env:"APPLICATION_PORT"`
	DebugMode       bool `env:"DEBUG" envDefault:"false"`
	AppReleaseMode  bool `env:"APP_RELEASE_MODE" envDefault:"false"`

	PGDatabaseHost              string        `env:"PG_DATABASE_HOST"`
	PGDatabasePort              int           `env:"PG_DATABASE_PORT"`
	PGDatabaseUsername          string        `env:"PG_DATABASE_USERNAME"`
	PGDatabasePassword          string        `env:"PG_DATABASE_PASSWORD"`
	PGDatabaseName              string        `env:"PG_DATABASE_NAME"`
	PGDatabaseConnectionTimeout time.Duration `env:"POSTGRES_DB_CONNECTION_TIMEOUT" envDefault:"180s"`
	PGTimeZone                  string        `env:"POSTGRES_TIME_ZONE"`

	PerPageItemLimit int `env:"PER_PAGE_ITEM_LIMIT" envDefault:"1000"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}

package config

import (
	"client/pkg/logging"
	"client/pkg/postgresql"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	REST    REST                 `envconfig:"REST"`
	Logging logging.LoggerConfig `envconfig:"LOG"`
}

type MigrationsConfig struct {
	Postgres postgresql.PostgreSQL `envconfig:"POSTGRES"`
}

// type RedisConfig

type REST struct {
	RESTHost string `envconfig:"REST_HOST" required:"false" default:"0.0.0.0"`
	RESTPort string `envconfig:"REST_PORT" required:"false" default:"8000"`
}

func NewFromEnv() *Config {
	cfg := &Config{}
	envconfig.MustProcess("", cfg)
	return cfg
}

func NewMigrationsFromEnv() *MigrationsConfig {
	c := MigrationsConfig{}
	envconfig.MustProcess("", &c)
	return &c
}

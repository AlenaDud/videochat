package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GRPC      Grpc                 `envconfig:"GRPC"`
	WebSocket string               `envconfig:"WEBSOCKET"`
	SFUClient string               `envconfig:"SFU_SERVICE"`
	Log       logging.LoggerConfig `envconfig:"LOG"`
}

type Grpc struct {
	GrpcHost string `envconfig:"GRPC_HOST" required:"true" default:"0.0.0.0"`
	GrpcPort string `envconfig:"GRPC_PORT" required:"true" default:"50002"`
}

type WebSocket struct {
	WebSocketHost string `envconfig:"WEBSOCKET_HOST" required:"true" default:"0.0.0.0"`
	WebSocketPort string `envconfig:"WEBSOCKET_PORT" required:"true" default:"3002"`
}
type SFUClient struct {
	SFUHost     string `envconfig:"USERS_HOST" required:"true" default:"0.0.0.0"`
	SFUGrpcPort string `envconfig:"USERS_GRPC_PORT" required:"true" default:"50000"`
}

func New() *Config {
	cfg := &Config{}
	envconfig.MustProcess("", cfg)
	return cfg
}

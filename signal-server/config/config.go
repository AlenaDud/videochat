package config

import (
	"github.com/kelseyhightower/envconfig"
	"signal-server/pkg/logging"
)

type Config struct {
	GRPC      Grpc                 `envconfig:"GRPC"`
	WebSocket WebSocket            `envconfig:"WEBSOCKET"`
	SFUClient SFUClient            `envconfig:"SFU_SERVICE"`
	Logging   logging.LoggerConfig `envconfig:"LOG"`
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
	SFUHost     string `envconfig:"SFU_HOST" required:"true" default:"0.0.0.0"`
	SFUGrpcPort string `envconfig:"SFU_GRPC_PORT" required:"true" default:"50000"`
}

func NewFromEnv() *Config {
	c := Config{}
	envconfig.MustProcess("", &c)
	return &c
}

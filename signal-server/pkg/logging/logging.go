package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type LoggerConfig struct {
	LogIndex  string `envconfig:"LOGGER_INDEX" required:"true" default:"todo"`
	IsDebug   bool   `envconfig:"DEBUG" required:"true" default:"false"`
	LogToFile bool   `envconfig:"LOG_TO_FILE" required:"true" default:"true"`
}

func NewLogger(cfg LoggerConfig) *zerolog.Logger {
	logger := log.With().Str("todo", cfg.LogIndex).Logger()

	if cfg.IsDebug {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	if cfg.LogToFile {
		file, err := os.Create("application.log")
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to create log file")
		}
		logger = logger.Output(file)
	} else {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	return &logger
}

package app

import (
	"client/config"
	"client/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	cfg           *config.Config
	logger        *zerolog.Logger
	router        *mux.Router
	clientService api.ClientService
}

func NewApp(
	cfg *config.Config,
) (*App, error) {
	logger := logging.NewLogger(cfg.Logging)

	clientService, err := api.NewClientService(cfg, logger)

	if err != nil {
		return nil, fmt.Errorf("[NewApp] create client service: %w", err)
	}

	return &App{
		cfg:           cfg,
		logger:        logger,
		clientService: clientService,
	}, nil
}

func (app *App) RunAPI() error {
	group := new(errgroup.Group)
	group.Go(func() error {
		err := rest.NewRestAPI(app.logger, app.clientService, app.cfg)
		if err != nil {
			return fmt.Errorf("[RunAPI] create rest api: %w", err)
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("[RunAPI] run: %w]", err)
	}
	return nil
}

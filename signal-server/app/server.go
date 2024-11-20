package app

import (
	"fmt"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"signal-server/config"
	"signal-server/internal/api/websocket"
	"signal-server/internal/services"
	"signal-server/pkg/logging"
)

type App struct {
	cfg           *config.Config
	logger        *zerolog.Logger
	signalService *services.SignalService
}

func NewApp(
	cfg *config.Config,
) (*App, error) {
	logger := logging.NewLogger(cfg.Logging)

	signalService := services.NewSignalService()

	return &App{
		cfg:           cfg,
		logger:        logger,
		signalService: signalService,
	}, nil
}

func (a *App) RunApp() error {
	group := new(errgroup.Group)
	group.Go(func() error {
		err := websocket.NewWebsocketAPI(a.cfg, a.logger, a.signalService)
		return fmt.Errorf("[RunApp] run WebSocketApi: %w", err)
	})
	//group.Go(func() error {
	//	err := grpc.NewGrpcApi(a.cfg, a.logger, a.signalService)
	//	return fmt.Errorf("[RunApp] run GrpcApi: %w", err)
	//})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("[RunApp] run %w", err)
	}
	return nil

}

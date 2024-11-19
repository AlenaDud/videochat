package app

import (
	"fmt"
	"signal-server/config"
	"signal-server/internal/api/grpc"
	"signal-server/internal/api/websocket"
	"signal-server/internal/clients/sfu"
	"signal-server/internal/services"
	"sync"
)

type App struct {
	cfg           *config.Config
	grpcServer    *grpc.Server
	wsServer      *websocket.Server
	signalService *services.SignalService
}

func NewApp(cfg *config.Config) (*App, error) {
	sfuClient, err := sfu.NewSFUClient(cfg.SFUServiceAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SFU: %w", err)
	}

	signalService := services.NewSignalService(sfuClient)
	wsServer := websocket.NewServer(signalService)
	grpcServer := grpc.NewServer(signalService)

	return &App{
		cfg:           cfg,
		grpcServer:    grpcServer,
		wsServer:      wsServer,
		signalService: signalService,
	}, nil
}

func (a *App) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := a.grpcServer.Start(a.cfg.GRPCPort); err != nil {
			fmt.Println("Failed to start gRPC server:", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.wsServer.Start(a.cfg.WebSocketPort); err != nil {
			fmt.Println("Failed to start WebSocket server:", err)
		}
	}()

	wg.Wait()
	return nil
}

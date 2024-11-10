package app

import (
	"net/http"
	"signal-server/internal/handlers"
	"signal-server/internal/services"

	"github.com/gorilla/mux"
)

type App struct {
	router           *mux.Router
	signalingService handlers.SignalingService
}

func NewApp() (*App, error) {
	signalingService := services.NewSignalingService(nil)

	return &App{
		signalingService: signalingService,
	}, nil
}

func (a *App) RunServer() {
	signalHandler := handlers.NewSignalHandler(a.signalingService)

	a.router = mux.NewRouter()
	a.router.HandleFunc("/ws", signalHandler.HandleWebSocket).Methods(http.MethodGet)
	http.ListenAndServe(":8080", a.router)
}

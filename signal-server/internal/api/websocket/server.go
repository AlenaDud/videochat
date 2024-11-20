package websocket

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"signal-server/config"
	"signal-server/internal/api"
)

func NewWebsocketAPI(
	cfg *config.Config,
	logger *zerolog.Logger,
	signalServer api.SignalService,
) error {
	signalHandler := NewSignalHandler(logger, signalServer)
	router := mux.NewRouter()

	router.HandleFunc("/ws/sdp/offer", signalHandler.handleSDPOffer)
	//router.HandleFunc("/ws/create", signalHandler.handleCreateRoom)
	//router.HandleFunc("/ws/join", signalHandler.handleJoinRoom)
	//router.HandleFunc("/ws/sdp/answer", signalHandler.handleSDPAnswer)
	//router.HandleFunc("/ws/ice-candidate", signalHandler.handleICECandidate)
	//router.HandleFunc("/ws/leave", signalHandler.handleLeaveRoom)
	//router.HandleFunc("/ws/chat", signalHandler.handleChatMessage)

	appAddr := fmt.Sprintf("%s:%s", cfg.WebSocket.WebSocketHost, cfg.WebSocket.WebSocketPort)
	logger.Info().Msgf("running websocket server at '%s'", appAddr)

	err := http.ListenAndServe(appAddr, router)

	if err != nil {
		return fmt.Errorf("[NewWebsocketAPI] listen ans serve: %w", err)
	}

	return nil
}

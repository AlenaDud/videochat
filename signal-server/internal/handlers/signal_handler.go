package handlers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"signal-server/internal/services"
)

type SignalHandler struct {
	signalingService services.SignalingService
	upgrader         websocket.Upgrader
}

func NewSignalHandler(signalingService services.SignalingService) *SignalHandler {
	return &SignalHandler{
		signalingService: signalingService,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (h *SignalHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	// Обработка сообщений WebSocket
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Обработка сигналов через сервисный слой
		h.signalingService.ProcessMessage(conn, messageType, message)
	}
}

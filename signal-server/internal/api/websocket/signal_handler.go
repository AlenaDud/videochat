package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"net/http"
	"signal-server/internal/api"
	"signal-server/internal/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type SignalHandler struct {
	logger       *zerolog.Logger
	signalServer api.SignalService
}

func NewSignalHandler(
	logger *zerolog.Logger,
	signalServer api.SignalService,
) *SignalHandler {
	return &SignalHandler{
		logger:       logger,
		signalServer: signalServer,
	}
}

func (h *SignalHandler) handleSDPOffer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to upgrade to WebSocket")
		return
	}
	defer conn.Close()

	for {
		var offer models.SDPMessage
		err := conn.ReadJSON(&offer)
		if err != nil {
			h.logger.Error().Msgf("[handleSDPOffer] read offer from ws: %v", err)
			break
		}
		h.logger.Info().Msgf("[handleSDPOffer] offer received: %v", offer)

		//answer, err := h.signalServer.SendOffer(offer) // Передача в SFU
		//if err != nil {
		//	h.logger.Error().Err(err).Msg("Failed to forward SDP offer to SFU")
		//	break
		//}
		//
		//err = conn.WriteJSON(answer) // Возврат SDP Answer клиенту
		//if err != nil {
		//	h.logger.Error().Err(err).Msg("Failed to send SDP answer to client")
		//	break
		//}
	}
}

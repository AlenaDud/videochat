package handler

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"webrtc-microservices/signal-server/pkg/grpc"
)

type SignalHandler struct {
	grpcClient grpc.SFUClient
}

func NewSignalHandler(grpcClient grpc.SFUClient) *SignalHandler {
	return &SignalHandler{grpcClient: grpcClient}
}

// WebSocket обработка сообщений от клиентов
func (h *SignalHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection error:", err)
		return
	}
	defer conn.Close()

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		switch msg["type"] {
		case "offer":
			h.handleOffer(conn, msg)
		case "answer":
			h.handleAnswer(conn, msg)
		case "ice-candidate":
			h.handleIceCandidate(conn, msg)
		default:
			log.Println("Unknown message type:", msg["type"])
		}
	}
}

// Обработка offer от клиента
func (h *SignalHandler) handleOffer(conn *websocket.Conn, msg map[string]interface{}) {
	sessionID := msg["session_id"].(string)
	userID := msg["user_id"].(string)
	offer := msg["offer"].(string)

	// Отправка offer на SFU
	response, err := h.grpcClient.SendOffer(sessionID, userID, offer)
	if err != nil {
		log.Println("Error sending offer to SFU:", err)
		return
	}

	// Отправка ответа клиенту
	conn.WriteJSON(map[string]string{"status": response.Status})
}

// Обработка answer от клиента
func (h *SignalHandler) handleAnswer(conn *websocket.Conn, msg map[string]interface{}) {
	sessionID := msg["session_id"].(string)
	userID := msg["user_id"].(string)
	answer := msg["answer"].(string)

	// Отправка answer на SFU
	response, err := h.grpcClient.SendAnswer(sessionID, userID, answer)
	if err != nil {
		log.Println("Error sending answer to SFU:", err)
		return
	}

	// Отправка ответа клиенту
	conn.WriteJSON(map[string]string{"status": response.Status})
}

// Обработка ICE кандидатов от клиента
func (h *SignalHandler) handleIceCandidate(conn *websocket.Conn, msg map[string]interface{}) {
	sessionID := msg["session_id"].(string)
	userID := msg["user_id"].(string)
	candidate := msg["candidate"].(string)

	// Отправка ICE кандидата на SFU
	response, err := h.grpcClient.SendIceCandidate(sessionID, userID, candidate)
	if err != nil {
		log.Println("Error sending ICE candidate to SFU:", err)
		return
	}

	// Отправка ответа клиенту
	conn.WriteJSON(map[string]string{"status": response.Status})
}

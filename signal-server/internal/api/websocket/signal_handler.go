package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"signal-server/internal/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Server struct {
	signalService *services.SignalService
	mu            sync.Mutex
}

func NewServer(signalService *services.SignalService) *Server {
	return &Server{signalService: signalService}
}

func (s *Server) Start(port string) error {
	http.HandleFunc("/ws", s.handleConnection)
	log.Println("WebSocket server started on port", port)
	return http.ListenAndServe(":"+port, nil)
}

func (s *Server) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Обработка сообщения и отправка в SFU
		go s.signalService.HandleSDP(message)
	}
}

package main

import (
	"log"
	"net/http"
	"signal-server/config"
	"signal-server/internal/handler"
	"signal-server/pkg/grpc"
)

func main() {
	// Инициализация конфигурации
	cfg := config.LoadConfig()

	// Инициализация gRPC клиента
	grpcClient := grpc.NewSFUClient(cfg.SFUAddress)

	// Инициализация обработчиков
	signalHandler := handler.NewSignalHandler(grpcClient)

	// Настройка WebSocket сервера
	http.HandleFunc("/ws", signalHandler.HandleWebSocket)

	// Запуск HTTP сервера
	log.Printf("Signal server started on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}

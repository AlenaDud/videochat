package main

import (
	"log"
	"signal-server/app"
)

func main() {
	// Инициализация приложения
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	// Запуск сервера WebSocket
	application.RunServer()
}

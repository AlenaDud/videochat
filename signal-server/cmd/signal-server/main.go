package main

import (
	"signal-server/app"
	"signal-server/config"
)

func main() {
	cfg := config.New()
	application, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}
	if err := application.Run(); err != nil {
		panic(err)
	}
}

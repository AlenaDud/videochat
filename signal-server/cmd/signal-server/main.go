package main

import (
	"signal-server/app"
	"signal-server/config"
)

func main() {
	cfg := config.NewFromEnv()

	appServer, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}

	err = appServer.RunApp()
	if err != nil {
		panic(err)
	}
}

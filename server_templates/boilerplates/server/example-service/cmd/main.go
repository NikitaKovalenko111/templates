package main

import (
	"example-service/internal/app"
	"example-service/internal/config"
	"os"
	"os/signal"
	"syscall"

	_ "example-service/docs"
)

//	@title			Example API
//	@version		0.1
//	@description	Example API Server

// BasePath /

func main() {
	cfg := config.MustLoad()

	app := app.New(cfg)

	go app.Run()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit

	app.Stop()
}

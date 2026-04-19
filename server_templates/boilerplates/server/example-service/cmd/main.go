package main

import (
	"example-service/internal/app"
	"example-service/internal/config"

	_ "example-service/docs"
)

//	@title			Example API
//	@version		0.1
//	@description	Example API Server

// BasePath /

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}

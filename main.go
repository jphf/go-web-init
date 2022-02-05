package main

import (
	"api/internal/util/config"
	"api/internal/util/logger"
	"api/server"
)

func main() {
	config.Setup()
	logger.Setup()

	s := server.CreateServer()
	s.Run()
}

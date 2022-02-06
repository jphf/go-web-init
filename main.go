package main

import (
	"github.com/jphf/go-web-init.git/internal/util/config"
	"github.com/jphf/go-web-init.git/internal/util/logger"
	"github.com/jphf/go-web-init.git/server"
)

func main() {
	config.Setup()
	logger.Setup()

	s := server.CreateServer()
	s.Run()
}

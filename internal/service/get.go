package service

import "github.com/jphf/go-web-init.git/internal/util/config"

type GetService struct{}

func (s GetService) GetName() string {
	return config.Config.Name
}

func (s GetService) GetUsage() string {
	return config.Config.Usage
}

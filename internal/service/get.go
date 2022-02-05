package service

import "api/internal/util/config"

func GetName() string {
	return config.Config.Name
}

func GetUsage() string {
	return config.Config.Usage
}

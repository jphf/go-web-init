package service

import (
	"api/internal/util/logger"
	"fmt"
)

func PostData(name, usage string) {
	log := logger.Logger

	log.Info(fmt.Sprintf("name=%s, usage=%s", name, usage))
}

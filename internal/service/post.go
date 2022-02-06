package service

import (
	"fmt"

	"github.com/jphf/go-web-init.git/internal/util/logger"
)

type PostService struct{}

func (s *PostService) PostData(name, usage string) {
	log := logger.Logger

	log.Info(fmt.Sprintf("name=%s, usage=%s", name, usage))
}

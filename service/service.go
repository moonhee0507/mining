package service

import (
	"mining/config"
	"mining/repository"

	"github.com/inconshreveable/log15"
)

type Service struct {
	config *config.Config

	repository *repository.Repository

	log log15.Logger
}

func NewService(config *config.Config, repository *repository.Repository) *Service {
	s := &Service {
		config: config,
		log: log15.New("module", "service"),
	}

	return s
}
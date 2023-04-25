package service

import "github.com/stefnef/Flowingo/m/internal/core/domain"

type InfoService struct {
}

func (service InfoService) getInfo() *domain.Info {
	return &domain.Info{
		Text: "Example Resource Server",
	}
}

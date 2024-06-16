package service

import "github.com/stefnef/Flowingo/m/internal/core/domain"

type InfoService interface {
	GetInfo() *domain.Info
}

type InfoServiceImpl struct {
}

func NewInfoService() InfoService {
	return &InfoServiceImpl{}
}

func (service InfoServiceImpl) GetInfo() *domain.Info {
	return &domain.Info{
		Text: "Example Resource Server",
	}
}

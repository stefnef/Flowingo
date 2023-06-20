package service

import "github.com/stefnef/Flowingo/m/internal/core/domain"

type ResourceService interface {
	GetResources() []domain.Resource
}

type ResourceServiceImpl struct {
}

func (r ResourceServiceImpl) GetResources() []domain.Resource {
	return []domain.Resource{
		{
			Id:          "some-id",
			Name:        "Some Name",
			MagicNumber: 41,
		},
	}
}

func NewResourceService() ResourceService {
	return &ResourceServiceImpl{}
}

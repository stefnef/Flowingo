package repository

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
)

type InternalResourceRepositoryImpl struct{}

var resourceData = []domain.Resource{
	{
		Id:          "some-id",
		Name:        "Some Name",
		MagicNumber: 41,
	},
	{
		Id:          "some-other-id",
		Name:        "Some Other Name",
		MagicNumber: 37,
	},
}

func (i InternalResourceRepositoryImpl) GetResources() []domain.Resource {
	return resourceData
}

func (i InternalResourceRepositoryImpl) GetResourceById(id string) (*domain.Resource, error) {
	for _, resource := range resourceData {
		if id == resource.Id {
			return &resource, nil
		}
	}
	return nil, domain.NotFoundError
}

func NewInternalResourceRepository() ResourceRepository {
	return &InternalResourceRepositoryImpl{}
}

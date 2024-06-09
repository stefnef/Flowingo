package repository

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"math/rand"
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

func NewInternalResourceRepository() ResourceRepository {
	return &InternalResourceRepositoryImpl{}
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
	return nil, domain.NewNotFoundError("resource", id)
}

func (i InternalResourceRepositoryImpl) ExistsResourceByName(name string) bool {
	for _, resource := range resourceData {
		if name == resource.Name {
			return true
		}
	}
	return false
}

func (i InternalResourceRepositoryImpl) SaveResource(name string) *domain.Resource {
	var magicNumber = rand.Int()

	return &domain.Resource{
		Id:          "",
		Name:        name,
		MagicNumber: magicNumber,
	}
}

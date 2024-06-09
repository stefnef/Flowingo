package repository

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/pkg"
	"math/rand"
)

type InternalResourceRepositoryImpl struct {
	generator pkg.Generator
}

var resourceData []domain.Resource

func NewInternalResourceRepository(generator pkg.Generator) ResourceRepository {
	resourceData = []domain.Resource{
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
	return &InternalResourceRepositoryImpl{
		generator: generator,
	}
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
	var magicNumber = rand.Int() //TODO move this to a generator service
	var id = i.generator.GenerateUUID()

	resource := &domain.Resource{
		Id:          id,
		Name:        name,
		MagicNumber: magicNumber,
	}

	resourceData = append(resourceData, *resource)

	return resource
}

package service

import (
	"fmt"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/repository"
)

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) (*domain.Resource, error)
}

type ResourceServiceImpl struct {
	resourceRepository repository.ResourceRepository
}

func (r *ResourceServiceImpl) GetResource(id string) (*domain.Resource, error) {
	var resource, err = r.resourceRepository.GetResourceById(id)

	if err != nil {
		panic(fmt.Errorf("%w: Resource with id '%s'", domain.NotFoundError, id))
	}

	return resource, nil
}

func (r *ResourceServiceImpl) GetResources() []domain.Resource {
	return r.resourceRepository.GetResources()
}

func NewResourceService(resourceRepository repository.ResourceRepository) ResourceService {
	return &ResourceServiceImpl{
		resourceRepository: resourceRepository,
	}
}

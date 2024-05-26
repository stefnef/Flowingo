package service

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/repository"
)

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) *domain.Resource
}

type ResourceServiceImpl struct {
	resourceRepository repository.ResourceRepository
}

func (r *ResourceServiceImpl) GetResource(id string) *domain.Resource {
	var resource, err = r.resourceRepository.GetResourceById(id)

	if err != nil {
		panic(err) //TODO catch panic
	}

	return resource
}

func (r *ResourceServiceImpl) GetResources() []domain.Resource {
	return r.resourceRepository.GetResources()
}

func NewResourceService(resourceRepository repository.ResourceRepository) ResourceService {
	return &ResourceServiceImpl{
		resourceRepository: resourceRepository,
	}
}

package service

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/repository"
)

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) (*domain.Resource, error)
	PostResource(resourceName string) (*domain.Resource, error)
}

type ResourceServiceImpl struct {
	resourceRepository repository.ResourceRepository
}

func (r *ResourceServiceImpl) GetResource(id string) (*domain.Resource, error) {
	var resource, err = r.resourceRepository.GetResourceById(id)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *ResourceServiceImpl) GetResources() []domain.Resource {
	return r.resourceRepository.GetResources()
}

func (r *ResourceServiceImpl) PostResource(resourceName string) (*domain.Resource, error) {
	if exists := r.resourceRepository.ExistsResourceByName(resourceName); exists != false {
		return nil, domain.NewAlreadyExistsError(resourceName)
	}

	return r.resourceRepository.SaveResource(resourceName), nil
}

func NewResourceService(resourceRepository repository.ResourceRepository) ResourceService {
	return &ResourceServiceImpl{
		resourceRepository: resourceRepository,
	}
}

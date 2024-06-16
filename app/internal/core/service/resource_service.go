package service

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/repository"
	"github.com/stefnef/Flowingo/m/pkg"
)

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) (*domain.Resource, error)
	PostResource(resourceName string) (*domain.Resource, error)
}

type ResourceServiceImpl struct {
	resourceRepository repository.ResourceRepository
	generator          pkg.Generator
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

	var magicNumber = r.generator.GenerateNumber()

	return r.resourceRepository.SaveResource(resourceName, magicNumber), nil
}

func NewResourceService(resourceRepository repository.ResourceRepository, generator pkg.Generator) ResourceService {
	return &ResourceServiceImpl{
		resourceRepository: resourceRepository,
		generator:          generator,
	}
}

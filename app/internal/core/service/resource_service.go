package service

import (
	"fmt"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
)

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) (*domain.Resource, error)
}

type ResourceServiceImpl struct {
}

var resources = []domain.Resource{
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

func (r *ResourceServiceImpl) GetResource(id string) (*domain.Resource, error) {
	for _, resource := range resources {
		if id == resource.Id {
			return &resource, nil
		}
	}
	//TODO handle notFound error
	return nil, fmt.Errorf("%w: Resource with id '%s'", domain.NotFoundError, id)
}

func (r *ResourceServiceImpl) GetResources() []domain.Resource {
	return resources
}

func NewResourceService() ResourceService {
	return &ResourceServiceImpl{}
}

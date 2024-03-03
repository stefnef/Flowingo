package service

import "github.com/stefnef/Flowingo/m/internal/core/domain"

type ResourceService interface {
	GetResources() []domain.Resource
	GetResource(id string) domain.Resource
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

func (r *ResourceServiceImpl) GetResource(id string) domain.Resource {
	for _, resource := range resources {
		if id == resource.Id {
			return resource
		}
	}
	//TODO handle notFound error
	return resources[0]
}

func (r *ResourceServiceImpl) GetResources() []domain.Resource {
	return resources
}

func NewResourceService() ResourceService {
	return &ResourceServiceImpl{}
}

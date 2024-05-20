package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResourceRepositoryMock struct{}

var getResources func() []domain.Resource
var getResourceById func(id string) (*domain.Resource, error)

func (repository *ResourceRepositoryMock) GetResources() []domain.Resource {
	return getResources()
}
func (repository *ResourceRepositoryMock) GetResourceById(id string) (*domain.Resource, error) {
	return getResourceById(id)
}

var resourceRepository = &ResourceRepositoryMock{}
var resourceService = service.NewResourceService(resourceRepository)

func TestGetResources(t *testing.T) {
	var expectedResources = []domain.Resource{
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
	getResources = func() []domain.Resource { return expectedResources }

	var resources = resourceService.GetResources()

	assert.NotNil(t, resources)
	assert.Equal(t, expectedResources, resources)
}

func TestGetResource(t *testing.T) {
	var expectedResource = domain.Resource{
		Id:          "some-id",
		Name:        "Something's Name",
		MagicNumber: 41,
	}
	getResourceById = func(id string) (*domain.Resource, error) {
		if id == "some-id" {
			return &expectedResource, nil
		} else {
			return nil, domain.NotFoundError
		}
	}

	var resource, _ = resourceService.GetResource("some-id")

	assert.NotNil(t, resource)
	assert.Equal(t, expectedResource, *resource)
}

func TestResourceServiceImpl_GetResource_throws(t *testing.T) {
	getResourceById = func(id string) (*domain.Resource, error) {
		return nil, domain.NotFoundError
	}

	assert.PanicsWithError(t, "not found: Resource with id 'i-do-not-exist'", func() {
		_, _ = resourceService.GetResource("i-do-not-exist")
	})
}

package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResourceRepositoryMock struct {
	Slots map[string]interface{}
}

var existsResourceByName func(name string) bool
var getResources func() []domain.Resource
var getResourceById func(id string) (*domain.Resource, error)
var saveResource func(name string, magicNumber int) *domain.Resource

func (repository *ResourceRepositoryMock) GetResources() []domain.Resource {
	return getResources()
}

func (repository *ResourceRepositoryMock) GetResourceById(id string) (*domain.Resource, error) {
	return getResourceById(id)
}

func (repository *ResourceRepositoryMock) ExistsResourceByName(name string) bool {
	return existsResourceByName(name)
}

func (repository *ResourceRepositoryMock) SaveResource(name string, magicNumber int) *domain.Resource {
	repository.Slots["SaveResource"] = []interface{}{name, magicNumber}
	return saveResource(name, magicNumber)
}

type generatorMock struct{}

func (g *generatorMock) GenerateUUID() string {
	return "some-uuid"
}

var magicNumber = 212121

func (g *generatorMock) GenerateNumber() int {
	return magicNumber
}

var resourceRepository = &ResourceRepositoryMock{Slots: make(map[string]interface{})}
var generator = &generatorMock{}
var resourceService = service.NewResourceService(resourceRepository, generator)

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
			return nil, domain.NewNotFoundError("resource", id)
		}
	}

	var resource, _ = resourceService.GetResource("some-id")

	assert.NotNil(t, resource)
	assert.Equal(t, expectedResource, *resource)
}

func TestResourceServiceImpl_GetResource_propagates_not_found_error(t *testing.T) {
	id := "i-do-not-exist"
	notFoundError := domain.NewNotFoundError("resource", id)

	getResourceById = func(id string) (*domain.Resource, error) {
		return nil, notFoundError
	}

	found, err := resourceService.GetResource(id)

	assert.Equal(t, notFoundError, err)
	assert.Nil(t, found)
}

func TestResourceServiceImpl_PostResource_throws_error_if_already_exists(t *testing.T) {
	name := "i-exist"
	var alreadyExistsError = domain.NewAlreadyExistsError(name)

	existsResourceByName = func(name string) bool {
		return true
	}

	resource, err := resourceService.PostResource(name)

	assert.Equal(t, alreadyExistsError, err)
	assert.Nil(t, resource)
}

func TestResourceServiceImpl_PostResource_saves(t *testing.T) {
	name := "new-resource"
	expectedResource := &domain.Resource{
		Id:          "some-id",
		Name:        name,
		MagicNumber: magicNumber,
	}

	existsResourceByName = func(name string) bool {
		return false
	}

	saveResource = func(name string, magicNumber int) *domain.Resource {
		return expectedResource
	}

	resource, err := resourceService.PostResource(name)

	assert.Nil(t, err)
	assert.NotNil(t, resource)
	assert.Equal(t, expectedResource, resource)

	var calls = resourceRepository.Slots["SaveResource"]
	assert.Equal(t, []interface{}{name, magicNumber}, calls)
}

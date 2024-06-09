package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

type GeneratorMock struct{}

func (g *GeneratorMock) GenerateUUID() string {
	return "some-uuid"
}

var generator = &GeneratorMock{}
var resourceRepository = repository.NewInternalResourceRepository(generator)
var resourceData []domain.Resource

func setup() {
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
	resourceRepository = repository.NewInternalResourceRepository(generator)
}

func TestResourceRepositoryImpl_Is_ResourceRepository(t *testing.T) {
	assert.Implements(t, (*repository.ResourceRepository)(nil), resourceRepository)
}

func TestResourceRepositoryImpl_GetResources(t *testing.T) {
	setup()

	var expectedResources = resourceData
	var resources = resourceRepository.GetResources()

	assert.NotNil(t, resources)
	assert.Equal(t, expectedResources, resources)
}

func TestResourceRepositoryImpl_GetResource(t *testing.T) {
	setup()

	var expectedResource = resourceData[0]
	var resource, _ = resourceRepository.GetResourceById("some-id")

	assert.NotNil(t, resource)
	assert.Equal(t, expectedResource, *resource)
}

func TestResourceRepositoryImpl_GetResource_By_id(t *testing.T) {
	setup()

	tests := []struct {
		id   string
		want domain.Resource
	}{
		{
			id:   "some-id",
			want: resourceData[0],
		},
		{
			id:   "some-other-id",
			want: resourceData[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			act, err := resourceRepository.GetResourceById(tt.id)
			assert.Equal(t, tt.want, *act)
			assert.Nil(t, err)
		})
	}
}

func TestResourceRepositoryImpl_GetResource_throws(t *testing.T) {
	setup()

	element, err := resourceRepository.GetResourceById("i-do-not-exist")
	assert.Nil(t, element)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "could not find resource 'resource' with id 'i-do-not-exist'")
}

func TestResourceRepositoryImpl_ExistsResourceByName(t *testing.T) {
	setup()

	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Some Name",
			want: true,
		},
		{
			name: "i-do-not-exist",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := resourceRepository.ExistsResourceByName(tt.name)
			assert.Equal(t, tt.want, act)
		})
	}
}

func TestResourceRepositoryImpl_SaveResource(t *testing.T) {
	setup()
	var expectedUUID = "some-uuid"

	resource := resourceRepository.SaveResource("some-new-name")

	assert.NotNil(t, resource)
	assert.Equal(t, expectedUUID, resource.Id)
	assert.Equal(t, "some-new-name", resource.Name)
	assert.NotEqual(t, 0, resource.MagicNumber)
}

func TestResourceRepositoryImpl_SaveResource_and_find_by_id(t *testing.T) {
	setup()

	savedResource := resourceRepository.SaveResource("new-resource")

	foundResource, err := resourceRepository.GetResourceById(savedResource.Id)

	assert.Nil(t, err)
	assert.NotNil(t, foundResource)
	assert.Equal(t, savedResource, foundResource)
}

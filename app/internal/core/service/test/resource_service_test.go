package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resourceService = service.NewResourceService()
var resourceData = []domain.Resource{
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

func TestGetResources(t *testing.T) {
	var expectedResources = resourceData
	var resources = resourceService.GetResources()

	assert.NotNil(t, resources)
	assert.Equal(t, expectedResources, resources)
}

func TestGetResource(t *testing.T) {
	var expectedResource = resourceData[0]
	var resource = resourceService.GetResource("some-id")

	assert.NotNil(t, resource)
	assert.Equal(t, expectedResource, resource)
}

func TestResourceServiceImpl_GetResource(t *testing.T) {
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
			act := resourceService.GetResource(tt.id)
			assert.Equal(t, tt.want, act)
		})
	}
}

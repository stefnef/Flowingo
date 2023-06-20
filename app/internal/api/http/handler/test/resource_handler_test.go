package test_test

import (
	"github.com/goccy/go-json"
	"github.com/stefnef/Flowingo/m/internal/api/http/dto"
	"github.com/stefnef/Flowingo/m/internal/api/http/handler"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type ResourceServiceMock struct {
}

var getResources func() []domain.Resource

func (service *ResourceServiceMock) GetResources() []domain.Resource {
	return getResources()
}

var (
	resourceService = &ResourceServiceMock{}
	resourceHandler = handler.NewResourceHandler(resourceService)
)

func TestResourceHandlerImpl_GetResources(t *testing.T) {
	var context, recorder = GetTestGinContext()
	var resourcesDto []dto.ResourceResponseDto
	var expectedDto = []dto.ResourceResponseDto{
		{Id: "resource-id", Name: "Name", MagicNumber: 21},
	}

	getResources = func() []domain.Resource {
		return []domain.Resource{
			{Id: "resource-id", Name: "Name", MagicNumber: 21},
		}
	}

	resourceHandler.GetResources(context)

	var err = json.Unmarshal(recorder.Body.Bytes(), &resourcesDto)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedDto, resourcesDto)
}

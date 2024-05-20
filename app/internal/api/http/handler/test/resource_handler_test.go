package handler_test

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
	slot *Slot
}

var getResources func() []domain.Resource
var getResource func() *domain.Resource

func (service *ResourceServiceMock) GetResources() []domain.Resource {
	return getResources()
}

const functionGetResource = "GetResource"

func (service *ResourceServiceMock) recordFunctionCall(function string, parameterName string, parameterValue string) {
	if service.slot != nil {
		service.slot.appendParameter(function, parameterName, parameterValue)
	}
}

func (service *ResourceServiceMock) GetResource(id string) *domain.Resource {
	service.recordFunctionCall(functionGetResource, "id", id)
	return getResource()
}

func initResourceHandlerSlot(t *testing.T) {
	resourceService.slot = &Slot{
		functions: map[string]*[]slotParam{
			functionGetResource: {},
		},
		t: t,
	}
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

func TestResourceHandlerImpl_GetResource_on_missing_id(t *testing.T) {
	initResourceHandlerSlot(t)
	var context, recorder = GetTestGinContext()

	resourceHandler.GetResource(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	resourceService.slot.verifyFunctionNotCalled(functionGetResource)
}

func doTestGetResource(t *testing.T, paramId string, expectedParamId string) {
	initResourceHandlerSlot(t)
	var context, recorder = GetTestGinContext()
	context.AddParam("id", paramId)
	var resourceDto dto.ResourceResponseDto
	var expectedDto = dto.ResourceResponseDto{
		Id:          expectedParamId,
		Name:        "Name",
		MagicNumber: 21,
	}

	getResource = func() *domain.Resource {
		return &domain.Resource{
			Id:          expectedParamId,
			Name:        "Name",
			MagicNumber: 21,
		}
	}

	resourceHandler.GetResource(context)

	var err = json.Unmarshal(recorder.Body.Bytes(), &resourceDto)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedDto, resourceDto)
	resourceService.slot.verify(functionGetResource, "id", expectedParamId)
}

func Test_GetResource(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		expectedId string
	}{
		{
			name:       "with normal parameter",
			id:         "resource-id",
			expectedId: "resource-id",
		},
		{
			name:       "it should trim id parameter",
			id:         " trimmed ",
			expectedId: "trimmed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doTestGetResource(t, tt.id, tt.expectedId)
		})
	}
}

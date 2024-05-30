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
var getResource func() (*domain.Resource, error)

func (service *ResourceServiceMock) GetResources() []domain.Resource {
	return getResources()
}

const functionGetResource = "GetResource"

func (service *ResourceServiceMock) recordFunctionCall(function string, parameterName string, parameterValue string) {
	if service.slot != nil {
		service.slot.appendParameter(function, parameterName, parameterValue)
	}
}

func (service *ResourceServiceMock) GetResource(id string) (*domain.Resource, error) {
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

func TestResourceHandlerImpl_GetResource_on_missing_parameter_id(t *testing.T) {
	initResourceHandlerSlot(t)
	var context, recorder = GetTestGinContext()

	resourceHandler.GetResource(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	resourceService.slot.verifyFunctionNotCalled(functionGetResource)
}

func TestResourceHandlerImpl_GetResource_error_on_missing_resource(t *testing.T) {
	initResourceHandlerSlot(t)
	var context, recorder = GetTestGinContext()
	id := "fake-id"
	context.AddParam("id", id)
	recorder.Code = 123
	notFoundError := domain.NewNotFoundError("fake", id)

	getResource = func() (*domain.Resource, error) {
		return nil, notFoundError
	}

	resourceHandler.GetResource(context)
	assert.NotEmpty(t, context.Errors)
	assert.Len(t, context.Errors, 1)
	assert.Equal(t, notFoundError, (*context.Errors[0]).Err)
	assert.Equal(t, 123, recorder.Code)
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

	getResource = func() (*domain.Resource, error) {
		return &domain.Resource{
			Id:          expectedParamId,
			Name:        "Name",
			MagicNumber: 21,
		}, nil
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

func TestResourceHandler_error_on_missing_param_name(t *testing.T) {
	var context, recorder = GetTestGinContext()
	recorder.Code = 000

	resourceHandler.PostResource(context)

	assert.NotEmpty(t, context.Errors)
	assert.Equal(t, 000, recorder.Code)
	assert.Len(t, context.Errors, 1)
	assert.Equal(t, "name is missing", (*context.Errors[0]).Err.Error())
}

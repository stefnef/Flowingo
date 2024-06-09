package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"net/http"
	"strings"
)

type ResourceHandler interface {
	GetResources(context *gin.Context)
	GetResource(context *gin.Context)
	PostResource(context *gin.Context)
}

type ResourceHandlerImpl struct {
	resourceService service.ResourceService
}

func (handler *ResourceHandlerImpl) GetResource(context *gin.Context) {
	id := strings.TrimSpace(context.Param("id"))
	if id == "" {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	resource, err := handler.resourceService.GetResource(id)
	if err != nil {
		_ = context.Error(err)
	} else {
		context.JSON(http.StatusOK, resource)
	}
}

func (handler *ResourceHandlerImpl) GetResources(context *gin.Context) {
	resources := handler.resourceService.GetResources()
	context.JSON(http.StatusOK, resources)
}

type resourceDto struct {
	Name string `json:"name" binding:"required"`
}

func (handler *ResourceHandlerImpl) PostResource(context *gin.Context) {
	var request *resourceDto
	err := context.BindJSON(&request)
	if err != nil {
		return
	}

	resource, err := handler.resourceService.PostResource(request.Name)

	if err != nil {
		_ = context.Error(err)
	} else {
		context.JSON(http.StatusCreated, resource)
	}
}

func NewResourceHandler(resourceService service.ResourceService) *ResourceHandlerImpl {
	return &ResourceHandlerImpl{
		resourceService: resourceService,
	}
}

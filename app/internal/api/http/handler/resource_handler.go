package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"net/http"
)

type ResourceHandler interface {
	GetResources(context *gin.Context)
}

type ResourceHandlerImpl struct {
	resourceService service.ResourceService
}

func NewResourceHandler(resourceService service.ResourceService) *ResourceHandlerImpl {
	return &ResourceHandlerImpl{
		resourceService: resourceService,
	}
}

func (handler *ResourceHandlerImpl) GetResources(c *gin.Context) {
	resources := handler.resourceService.GetResources()
	c.JSON(http.StatusOK, resources)
}

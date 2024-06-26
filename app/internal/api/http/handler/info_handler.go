package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"net/http"
)

type InfoHandler interface {
	GetInfo(c *gin.Context)
}

type InfoHandlerImpl struct {
	infoService service.InfoService
}

func NewInfoHandler(infoService service.InfoService) *InfoHandlerImpl {
	return &InfoHandlerImpl{
		infoService: infoService,
	}
}

func (handler *InfoHandlerImpl) GetInfo(context *gin.Context) {
	info := handler.infoService.GetInfo()
	context.JSON(http.StatusOK, info)
}

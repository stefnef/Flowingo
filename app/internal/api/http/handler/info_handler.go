package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"net/http"
)

type InfoHandler struct {
	infoService service.InfoService
}

func NewInfoHandler(infoService service.InfoService) *InfoHandler {
	return &InfoHandler{
		infoService: infoService,
	}
}

func (handler *InfoHandler) GetInfo(c *gin.Context) {
	info := handler.infoService.GetInfo()
	c.JSON(http.StatusOK, info)
}

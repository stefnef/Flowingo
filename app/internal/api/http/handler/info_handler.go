package handler

import (
	"github.com/stefnef/Flowingo/m/internal/api/http/dto"
	"github.com/stefnef/Flowingo/m/internal/core/service"
)

type InfoHandler struct {
	infoService service.InfoService
}

func (handler *InfoHandler) getInfo() *dto.InfoDto {
	info := handler.infoService.GetInfo()
	return &dto.InfoDto{Text: string(info.Text)}
}

func NewInfoHandler(infoService service.InfoService) *InfoHandler {
	return &InfoHandler{
		infoService: infoService,
	}
}

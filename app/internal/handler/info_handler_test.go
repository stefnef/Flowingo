package handler

import (
	"github.com/stefnef/Flowingo/m/api"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"testing"
)

type InfoServiceMock struct {
}

var getInfo func() *domain.Info

func (service *InfoServiceMock) getInfo() *domain.Info {
	return getInfo()
}

type InfoHandler struct {
}

func (handler *InfoHandler) getInfo() *api.InfoDto {
	return &api.InfoDto{Text: ""}
}

var handler = &InfoHandler{}

func TestGetInfo(t *testing.T) {
	getInfo = func() *domain.Info {
		return &domain.Info{Text: "Text"}
	}

	var infoDto = handler.getInfo()

	if infoDto == nil {
		t.Fatalf("info dto not created")
	}
}

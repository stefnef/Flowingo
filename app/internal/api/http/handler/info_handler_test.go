package handler

import (
	"github.com/stefnef/Flowingo/m/internal/api/http/dto"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"testing"
)

type InfoServiceMock struct {
}

var getInfo func() *domain.Info

func (service *InfoServiceMock) GetInfo() *domain.Info {
	return getInfo()
}

var infoService = &InfoServiceMock{}
var handler = NewInfoHandler(infoService)

func TestGetInfo(t *testing.T) {
	expectedDto := &dto.InfoDto{
		Text: "Text",
	}
	getInfo = func() *domain.Info {
		return &domain.Info{Text: "Text"}
	}

	var infoDto = handler.getInfo()

	assertThatIsNotNil(infoDto, t)
	assertThatIsEqualTo(infoDto, expectedDto, t)
}

func assertThatIsNotNil(act *dto.InfoDto, t *testing.T) {
	if act == nil {
		t.Fatalf("info dto not created")
	}
}

func assertThatIsEqualTo(act *dto.InfoDto, exp *dto.InfoDto, t *testing.T) {
	if *act != *exp {
		t.Fatalf("Unequal: act '%s', exp: '%s'", *act, *exp)
	}
}

package handler

import (
	"github.com/goccy/go-json"
	"github.com/stefnef/Flowingo/m/internal/api/http/dto"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
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
	var context, recorder = GetTestGinContext()
	var infoDto *dto.InfoDto
	var expectedDto = &dto.InfoDto{
		Text: "Text",
	}

	getInfo = func() *domain.Info {
		return &domain.Info{Text: "Text"}
	}

	handler.GetInfo(context)
	var err = json.Unmarshal(recorder.Body.Bytes(), &infoDto)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedDto, infoDto)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type infoHandlerMock struct {
}

func (infoHandler *infoHandlerMock) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &response)
}

type responseMock struct {
	Text string `json:"Text"`
}

var response responseMock
var infoHandler = infoHandlerMock{}

var router = NewRouter(&infoHandler)

func TestRouters(t *testing.T) {
	var infoResponse responseMock
	response = responseMock{
		Text: "getInfo",
	}

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)
	router.ServeHTTP(recorder, req)

	err := json.Unmarshal(recorder.Body.Bytes(), &infoResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, infoResponse)
	assert.Equal(t, "getInfo", infoResponse.Text)
}

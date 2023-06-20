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
	ctx.JSON(http.StatusOK, &responseMock{
		Text: "getInfo",
	})
}

type responseMock struct {
	Text string `json:"Text"`
}

func TestRouters(t *testing.T) {
	gin.SetMode(gin.TestMode)

	infoHandler := infoHandlerMock{}
	router := NewRouter(&infoHandler)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)

	var infoResponse responseMock

	router.ServeHTTP(recorder, req)

	err := json.Unmarshal(recorder.Body.Bytes(), &infoResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, infoResponse)
	assert.Equal(t, "getInfo", infoResponse.Text)
}

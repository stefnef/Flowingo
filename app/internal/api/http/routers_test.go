package http

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type responseMock struct {
	Text string `json:"Text"`
}

var response responseMock

type infoHandlerMock struct{}
type resourceHandlerMock struct{}

func createResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &response)
}

func (infoHandler *infoHandlerMock) GetInfo(ctx *gin.Context) {
	createResponse(ctx)
}

func (resourceHandler *resourceHandlerMock) GetResources(ctx *gin.Context) {
	createResponse(ctx)
}

func (resourceHandler *resourceHandlerMock) GetResource(ctx *gin.Context) {
	createResponse(ctx)
}

var infoHandler = infoHandlerMock{}
var resourceHandler = resourceHandlerMock{}

var router = NewRouter(&infoHandler, &resourceHandler)

func TestInfoGet(t *testing.T) {
	var infoResponse responseMock
	response = responseMock{
		Text: "getInfo",
	}

	recorder := doRequest("GET", "/info")

	err := json.Unmarshal(recorder.Body.Bytes(), &infoResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, infoResponse)
	assert.Equal(t, "getInfo", infoResponse.Text)
}

func TestResourceGet(t *testing.T) {
	var getAllResourcesResponse responseMock
	response = responseMock{
		Text: "getAllResources",
	}

	recorder := doRequest("GET", "/resource")

	err := json.Unmarshal(recorder.Body.Bytes(), &getAllResourcesResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, getAllResourcesResponse)
	assert.Equal(t, "getAllResources", getAllResourcesResponse.Text)
}

func TestResourceGetSingle(t *testing.T) {
	var getSingleResourcesResponse responseMock
	response = responseMock{
		Text: "getResource",
	}

	recorder := doRequest("GET", "/resource/some-id")

	err := json.Unmarshal(recorder.Body.Bytes(), &getSingleResourcesResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, getSingleResourcesResponse)
	assert.Equal(t, "getResource", getSingleResourcesResponse.Text)
}

func doRequest(method string, url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)
	router.ServeHTTP(recorder, req)
	return recorder
}

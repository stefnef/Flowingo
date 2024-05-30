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
type errorHandlerMock struct{}

func (infoHandler *infoHandlerMock) GetInfo(ctx *gin.Context) {
	response.Text += "GetInfo"
	ctx.JSON(http.StatusOK, &response)
}

func (resourceHandler *resourceHandlerMock) GetResources(ctx *gin.Context) {
	response.Text += "GetResources"
	ctx.JSON(http.StatusOK, &response)
}

func (resourceHandler *resourceHandlerMock) GetResource(ctx *gin.Context) {
	response.Text += "ResourceGetSingle"
	ctx.JSON(http.StatusOK, &response)
}

func (resourceHandler *resourceHandlerMock) PostResource(ctx *gin.Context) {
	response.Text += "ResourcePost"
	ctx.JSON(http.StatusOK, &response)
}

func (errorHandler *errorHandlerMock) HandleErrors(ctx *gin.Context) {
	response.Text = "errorHandling -> "
	ctx.Next()
}

var (
	infoHandler     = infoHandlerMock{}
	resourceHandler = resourceHandlerMock{}
	errorHandler    = errorHandlerMock{}
)

var router = NewRouter(&infoHandler, &resourceHandler, &errorHandler)

func setup() {
	response = responseMock{Text: ""}
}

func TestInfoGet(t *testing.T) {
	setup()
	var infoResponse responseMock

	recorder := doRequest("GET", "/info")

	err := json.Unmarshal(recorder.Body.Bytes(), &infoResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, infoResponse)
	assert.Equal(t, "errorHandling -> GetInfo", infoResponse.Text)
}

func TestResourceGet(t *testing.T) {
	setup()
	var getAllResourcesResponse responseMock

	recorder := doRequest("GET", "/resource")

	err := json.Unmarshal(recorder.Body.Bytes(), &getAllResourcesResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, getAllResourcesResponse)
	assert.Equal(t, "errorHandling -> GetResources", getAllResourcesResponse.Text)
}

func TestResourceGetSingle(t *testing.T) {
	setup()
	var getSingleResourcesResponse responseMock
	recorder := doRequest("GET", "/resource/some-id")

	err := json.Unmarshal(recorder.Body.Bytes(), &getSingleResourcesResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, getSingleResourcesResponse)
	assert.Equal(t, "errorHandling -> ResourceGetSingle", getSingleResourcesResponse.Text)
}

func TestResourcePost(t *testing.T) {
	setup()
	var postSingleResourceResponse responseMock
	recorder := doRequest("POST", "/resource")

	err := json.Unmarshal(recorder.Body.Bytes(), &postSingleResourceResponse)

	assert.Nil(t, err)
	assert.NotEmpty(t, postSingleResourceResponse)
	assert.Equal(t, "ResourcePost", postSingleResourceResponse.Text)
}

func doRequest(method string, url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)
	router.ServeHTTP(recorder, req)
	return recorder
}

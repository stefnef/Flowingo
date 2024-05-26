package handler_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func GetTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := createContextAndEngine(w)

	return ctx, w
}

func GetTestGinEngine() (*gin.Context, *httptest.ResponseRecorder, *gin.Engine) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, engine := createContextAndEngine(w)

	return ctx, w, engine
}

func createContextAndEngine(w *httptest.ResponseRecorder) (*gin.Context, *gin.Engine) {
	ctx, engine := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	return ctx, engine
}

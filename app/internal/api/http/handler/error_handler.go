package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorHandler interface {
	HandleErrors(context *gin.Context)
}

type ErrorHandlerImpl struct {
}

func (errorHandler *ErrorHandlerImpl) HandleErrors(context *gin.Context) {
	for _, _ = range context.Errors {
		context.AbortWithStatus(http.StatusNotFound)
	}
}

func NewErrorHandler() *ErrorHandlerImpl {
	return &ErrorHandlerImpl{}
}

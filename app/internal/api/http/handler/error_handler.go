package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	HandleErrors(context *gin.Context)
}

type ErrorHandlerImpl struct {
}

func (errorHandler *ErrorHandlerImpl) HandleErrors(ctx *gin.Context) {
	context.TODO()
}

func NewErrorHandler() *ErrorHandlerImpl {
	return &ErrorHandlerImpl{}
}

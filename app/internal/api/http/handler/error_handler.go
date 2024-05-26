package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"net/http"
)

type ErrorHandler interface {
	HandleErrors(context *gin.Context)
}

type ErrorHandlerImpl struct {
}

func NewErrorHandler() *ErrorHandlerImpl {
	return &ErrorHandlerImpl{}
}

func (errorHandler *ErrorHandlerImpl) HandleErrors(context *gin.Context) {
	context.Next()

	errorHandler.handleError(context)
}

func (errorHandler *ErrorHandlerImpl) handleError(context *gin.Context) {
	for _, err := range context.Errors {
		var notFoundError *domain.NotFoundError
		switch {
		case errors.As(err.Err, &notFoundError):
			context.AbortWithStatusJSON(http.StatusNotFound, err.JSON())
		default:
			internalServerError := map[string]string{"error": "internal server error"}
			context.AbortWithStatusJSON(http.StatusInternalServerError, internalServerError)
		}
	}
}

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

func (errorHandler *ErrorHandlerImpl) HandleErrors(context *gin.Context) {
	for _, err := range context.Errors {
		var notFoundError *domain.NotFoundError
		switch {
		case errors.As(err.Err, &notFoundError):
			context.AbortWithStatusJSON(http.StatusNotFound, err.JSON())
		default:
			context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
	}
}

func NewErrorHandler() *ErrorHandlerImpl {
	return &ErrorHandlerImpl{}
}

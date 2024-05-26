package handler_test

import (
	"github.com/stefnef/Flowingo/m/internal/api/http/handler"
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var errorHandler = handler.NewErrorHandler()

func Test_should_implement_interface(t *testing.T) {
	assert.Implements(t, (*handler.ErrorHandler)(nil), errorHandler)
}

func Test_should_do_nothing_if_there_is_no_error(t *testing.T) {
	var context, recorder = GetTestGinContext()
	recorder.Code = http.StatusCreated

	errorHandler.HandleErrors(context)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func Test_should_map_not_found(t *testing.T) {
	var context, recorder = GetTestGinContext()
	recorder.Code = http.StatusCreated

	var err = domain.NewNotFoundError("fake", "some-id")
	_ = context.Error(err)

	errorHandler.HandleErrors(context)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

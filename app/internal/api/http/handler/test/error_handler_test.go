package handler_test

import (
	"errors"
	"github.com/goccy/go-json"
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
	var errorDto map[string]string
	_ = context.Error(err)

	errorHandler.HandleErrors(context)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	var errJson = json.Unmarshal(recorder.Body.Bytes(), &errorDto)

	assert.Nil(t, errJson)
	assert.Equal(t, map[string]string{"error": "could not find resource 'fake' with id 'some-id'"}, errorDto)
}

func Test_should_unknown_error_to_server_error(t *testing.T) {
	var context, recorder = GetTestGinContext()
	recorder.Code = http.StatusCreated
	var err = errors.New("i am an unknown error")
	var errorDto map[string]string
	_ = context.Error(err)

	errorHandler.HandleErrors(context)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	var errJson = json.Unmarshal(recorder.Body.Bytes(), &errorDto)

	assert.Nil(t, errJson)
	assert.Equal(t, map[string]string{"error": "internal server error"}, errorDto)
}

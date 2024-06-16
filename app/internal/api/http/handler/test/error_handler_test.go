package handler_test

import (
	"errors"
	"github.com/gin-gonic/gin"
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

func Test_it_should_call_next_before_error_handling(t *testing.T) {
	var _, recorder, engine = GetTestGinEngine()
	var getHandler = func(context *gin.Context) {
		_ = context.Error(errors.New("here is some error"))
	}
	engine.GET("/", errorHandler.HandleErrors, getHandler)
	req, _ := http.NewRequest("GET", "/", nil)

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func Test_should_do_nothing_if_there_is_no_error(t *testing.T) {
	var context, recorder = GetTestGinContext()
	recorder.Code = http.StatusCreated

	errorHandler.HandleErrors(context)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func Test_should_map_error(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		status  int
		message string
	}{
		{
			"not found",
			domain.NewNotFoundError("fake", "some-id"),
			http.StatusNotFound,
			"could not find resource 'fake' with id 'some-id'",
		},

		{
			"already exists",
			domain.NewAlreadyExistsError("fake"),
			http.StatusBadRequest,
			"resource with name 'fake' already exists",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var context, recorder = GetTestGinContext()
			recorder.Code = http.StatusCreated
			var errorDto map[string]string
			_ = context.Error(test.err)

			errorHandler.HandleErrors(context)

			assert.Equal(t, test.status, recorder.Code)
			var errJson = json.Unmarshal(recorder.Body.Bytes(), &errorDto)

			assert.Nil(t, errJson)
			assert.Equal(t, map[string]string{"error": test.message}, errorDto)
		})
	}
}

func Test_should_map_unknown_error_to_server_error(t *testing.T) {
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

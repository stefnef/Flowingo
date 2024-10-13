package handler_test

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stefnef/Flowingo/m/internal/api/http/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type AuthServiceMock struct{}

var verifyAuth func(token string) error

func (authServiceMock AuthServiceMock) VerifyAuth(token string) error {
	return verifyAuth(token)
}

var authMockService = &AuthServiceMock{}
var authHandler = handler.NewAuthHandler(authMockService)

func Test_should_implement_auth_interface(t *testing.T) {
	assert.Implements(t, (*handler.AuthHandler)(nil), authHandler)
}

func Test_it_should_not_call_next_handler_if_there_is_an_error(t *testing.T) {
	var _, recorder, engine = GetTestGinEngine()
	var someHandler = func(context *gin.Context) {
		context.AbortWithStatus(400)
	}
	verifyAuth = func(token string) error {
		return errors.New("auth error")
	}
	engine.GET("/", authHandler.VerifyAuthenticated, someHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, 403, recorder.Code)
}

func Test_it_should_not_call_next_handler_if_there_is_no_error(t *testing.T) {
	var _, recorder, engine = GetTestGinEngine()
	var someHandler = func(context *gin.Context) {
		context.AbortWithStatus(200)
	}
	verifyAuth = func(token string) error {
		return nil
	}
	engine.GET("/", authHandler.VerifyAuthenticated, someHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code)
}

func Test_it_should_extract_token_from_header(t *testing.T) {
	var _, recorder, engine = GetTestGinEngine()

	verifyAuth = func(token string) error {
		return nil
	}
	engine.GET("/", authHandler.VerifyAuthenticated)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code)
}

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
		if token != "THE_TOKEN" {
			return errors.New("WRONG token")
		}
		return nil
	}

	engine.GET("/", authHandler.VerifyAuthenticated)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "THE_TOKEN"))

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code)
}

func Test_it_should_throw_an_error_if_token_is_not_present_or_in_wrong_format(t *testing.T) {
	var _, recorder, engine = GetTestGinEngine()

	verifyAuth = func(token string) error {
		if token != "THE_TOKEN" {
			return errors.New("WRONG token")
		}
		return nil
	}

	engine.GET("/", authHandler.VerifyAuthenticated)
	req, _ := http.NewRequest("GET", "/", nil)

	tests := []struct {
		id            string
		authorization string
		token         string
	}{
		{
			id:            "empty strings",
			authorization: "",
			token:         "",
		},
		{
			id:            "bearer missing",
			authorization: "Authorization",
			token:         fmt.Sprintf("Something Wrong %s", "THE_TOKEN"),
		},
		{
			id:            "missing space after bearer",
			authorization: "Authorization",
			token:         fmt.Sprintf("Bearer%s", "THE_TOKEN"),
		},
		{
			id:            "missing Authorization header",
			authorization: "Some-Other",
			token:         fmt.Sprintf("Bearer %s", "THE_TOKEN"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			req.Header.Add(tt.authorization, tt.token)
			engine.ServeHTTP(recorder, req)
			assert.Equal(t, 403, recorder.Code)
		})
	}
}

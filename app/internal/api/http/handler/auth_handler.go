package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	VerifyAuthenticated(context *gin.Context)
}

type AuthService interface {
	VerifyAuth(token string) error
}

type AuthHandlerImpl struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandlerImpl {
	return &AuthHandlerImpl{authService: authService}
}

func (authHandler *AuthHandlerImpl) VerifyAuthenticated(context *gin.Context) {
	var err = authHandler.authService.VerifyAuth("some-token") //TODO get token from context
	if err != nil {
		context.AbortWithStatus(403)
	}
	context.Next()
}

package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
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
	var authHeader = context.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	//TODO check not empty before giving to service

	var err = authHandler.authService.VerifyAuth(tokenString)
	if err != nil {
		context.AbortWithStatus(403)
	}
	context.Next()
}

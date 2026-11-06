package view

import (
	"time"

	"github.com/gin-gonic/gin"
)

func AuthErrorView(error string) gin.H {
	return gin.H{
		"error": error,
	}
}

func AuthSuccessView(token string, expiresIn time.Duration) gin.H {
	return gin.H{
		"AccessToken": token,
		"TokenType":   "Bearer",
		"ExpiresIn":   expiresIn,
	}
}

package view

import "github.com/gin-gonic/gin"

func ErrorView(error string) gin.H {
	return gin.H{
		"error": error,
	}
}

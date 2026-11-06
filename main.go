package main

import (
	"github/giovanifz/todolist-go/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/login", controllers.Login)

	router.Run("localhost:8080")
}

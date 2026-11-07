package main

import (
	"github/giovanifz/todolist-go/controllers"
	connection "github/giovanifz/todolist-go/model/connection"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		panic("Error when loading .env")
	}

	connection.ConnectDatabase()

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/login", controllers.Login)

	router.Run("localhost:8080")
}

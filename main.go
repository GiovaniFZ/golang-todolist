package main

import (
	"github/giovanifz/todolist-go/controllers"
	"github/giovanifz/todolist-go/middlewares"
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

	db := connection.ConnectDatabase()

	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	router.POST("/sign-up", func(c *gin.Context) {
		controllers.SignUp(c, db)
	})
	router.Use(middlewares.GinJWTMiddleware())
	{
		router.GET("/tasks", controllers.GetTasks)
	}

	router.Run("localhost:8080")
}

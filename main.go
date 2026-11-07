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

	db := connection.ConnectDatabase()

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	router.POST("/sign-up", func(c *gin.Context) {
		controllers.SignUp(c, db)
	})

	router.Run("localhost:8080")
}

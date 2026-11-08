package main

import (
	"github/giovanifz/todolist-go/controllers"
	"github/giovanifz/todolist-go/middlewares"
	connection "github/giovanifz/todolist-go/model/connection"
	"github/giovanifz/todolist-go/model/repositories/gorm"

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

	userRepository := gorm.NewGormUserRepository(db)
	taskRepository := gorm.NewGormTaskRepository(db)

router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, userRepository)
	})
	router.POST("/sign-up", func(c *gin.Context) {
		controllers.SignUp(c, userRepository)
	})
	router.Use(middlewares.GinJWTMiddleware())
	{
		router.GET("/tasks", func(c *gin.Context) {
			controllers.GetUserTasks(c, userRepository, taskRepository)
		})

		router.POST("/tasks", func(c *gin.Context) {
			controllers.CreateUserTask(c, userRepository, taskRepository)
		})
	}

	router.Run("localhost:8080")
}

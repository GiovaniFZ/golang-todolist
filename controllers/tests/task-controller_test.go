package test

import (
	"github/giovanifz/todolist-go/controllers"
	in_memory "github/giovanifz/todolist-go/model/repositories/in-memory"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUserTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	taskRepository := in_memory.NewInMemoryTaskRepository()
	userRepository := in_memory.NewInMemoryUserRepository()

	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/tasks", nil)

	controllers.CreateUserTask(c, userRepository, taskRepository)
}

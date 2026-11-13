package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github/giovanifz/todolist-go/controllers"
	"github/giovanifz/todolist-go/controllers/mocks"
	models "github/giovanifz/todolist-go/model"
	in_memory "github/giovanifz/todolist-go/model/repositories/in-memory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestCreateUserTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	taskRepository := in_memory.NewInMemoryTaskRepository()
	userRepository := in_memory.NewInMemoryUserRepository()

	body, err := json.Marshal(mocks.MockTask)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := gin.CreateTestContext(w)

	password_hash := controllers.HashPassword(mocks.MockLoginRequest.Password)
	user := models.User{
		Name:          mocks.MockRegisterRequest.Name,
		Email:         mocks.MockRegisterRequest.Email,
		Role:          "user",
		Tasks:         []models.Task{},
		Password_Hash: password_hash,
	}

	userRepository.Create(
		context.Background(),
		&user,
	)

	claims := &models.Claims{
		Subject: user.Id,
	}
	c.Set("userClaims", claims)

	c.Request = httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		bytes.NewReader(body),
	)

	controllers.CreateUserTask(c, userRepository, taskRepository)

	if c.Writer.Status() != http.StatusCreated {
		t.Errorf("Expected status %d, received status %d", http.StatusCreated, c.Writer.Status())
	}
}

func TestUpdateUserTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	taskRepository := in_memory.NewInMemoryTaskRepository()
	userRepository := in_memory.NewInMemoryUserRepository()

	body, err := json.Marshal(mocks.MockTask)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := gin.CreateTestContext(w)
	password_hash := controllers.HashPassword(mocks.MockLoginRequest.Password)
	user := models.User{
		Name:          mocks.MockRegisterRequest.Name,
		Email:         mocks.MockRegisterRequest.Email,
		Role:          "user",
		Tasks:         []models.Task{},
		Password_Hash: password_hash,
	}

	taskId := uuid.NewString()
	task := models.Task{
		ID:          taskId,
		Name:        mocks.MockTask.Name,
		Description: mocks.MockTask.Description,
		UserID:      user.Id,
	}
	taskRepository.Create(
		context.Background(),
		&task,
	)

	userRepository.Create(
		context.Background(),
		&user,
	)
	claims := &models.Claims{
		Subject: user.Id,
	}
	c.Set("userClaims", claims)

	c.Request = httptest.NewRequest(
		http.MethodPatch,
		"/tasks/",
		bytes.NewReader(body),
	)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: taskId})
	controllers.UpdateUserTask(c, userRepository, taskRepository)

	bodyToString := w.Body.String()

	fmt.Println(bodyToString)
	fmt.Println("/tasks/" + taskId)

	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status %d, received status %d", http.StatusOK, c.Writer.Status())
	}
}

func TestDeleteUserTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	taskRepository := in_memory.NewInMemoryTaskRepository()
	userRepository := in_memory.NewInMemoryUserRepository()

	c, _ := gin.CreateTestContext(w)
	password_hash := controllers.HashPassword(mocks.MockLoginRequest.Password)
	user := models.User{
		Name:          mocks.MockRegisterRequest.Name,
		Email:         mocks.MockRegisterRequest.Email,
		Role:          "user",
		Tasks:         []models.Task{},
		Password_Hash: password_hash,
	}

	taskId := uuid.NewString()
	task := models.Task{
		ID:          taskId,
		Name:        mocks.MockTask.Name,
		Description: mocks.MockTask.Description,
		UserID:      user.Id,
	}
	taskRepository.Create(
		context.Background(),
		&task,
	)

	userRepository.Create(
		context.Background(),
		&user,
	)
	claims := &models.Claims{
		Subject: user.Id,
	}
	c.Set("userClaims", claims)

	c.Request = httptest.NewRequest(
		http.MethodDelete,
		"/tasks/",
		nil,
	)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: taskId})
	controllers.DeleteUserTask(c, userRepository, taskRepository)

	if c.Writer.Status() != http.StatusNoContent {
		t.Errorf("Expected status %d, received status %d", http.StatusOK, c.Writer.Status())
	}
}

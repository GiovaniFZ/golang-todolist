package controllers

import (
	"context"
	"encoding/json"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/model/repositories"
	"github/giovanifz/todolist-go/view"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func GetUserTasks(c *gin.Context, userRepository repositories.UserRepository, taskRepository repositories.TaskRepository) {
	ctx := context.Background()
	claim, exists := c.Get("userClaims")
	if !exists {
		c.JSON(404, view.ErrorView("userClaims not found"))
		return
	}

	user, err := userRepository.FindByID(ctx, claim.(*models.Claims).Subject)
	if err != nil {
		c.JSON(404, view.ErrorView("User not found"))
		return
	}

	tasks, err := taskRepository.FindByUserID(ctx, user.Id)
	if err != nil {
		c.JSON(404, view.ErrorView("User tasks not found"))
		return
	}

	c.JSON(http.StatusOK, view.TaskView(tasks))
}

func CreateUserTask(c *gin.Context, userRepository repositories.UserRepository, taskRepository repositories.TaskRepository) {
	ctx := context.Background()
	claim, exists := c.Get("userClaims")
	if !exists {
		c.JSON(404, view.ErrorView("userClaims not found"))
		return
	}

	var req models.TaskRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView("Invalid body request"))
		return
	}

	if req.Name == "" || req.Description == "" || req.StartsAt.IsZero() || req.EndsAt.IsZero() {
		c.JSON(http.StatusBadRequest, view.ErrorView("Invalid task data"))
		return
	}

	user, err := userRepository.FindByID(ctx, claim.(*models.Claims).Subject)
	if err != nil {
		c.JSON(404, view.ErrorView("User not found"))
		return
	}

	newTaskId := uuid.NewString()
	newTask := &models.Task{
		ID:          newTaskId,
		UserId:      user.Id,
		Name:        req.Name,
		Description: req.Description,
		StartsAt:    req.StartsAt,
		EndsAt:      req.EndsAt,
	}
	taskRepository.Create(ctx, newTask)

	c.Status(http.StatusCreated)
}

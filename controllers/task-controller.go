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
		UserID:      user.Id,
		Name:        req.Name,
		Description: req.Description,
		StartsAt:    req.StartsAt,
		EndsAt:      req.EndsAt,
	}
	taskRepository.Create(ctx, newTask)

	c.Status(http.StatusCreated)
}

func UpdateUserTask(c *gin.Context, userRepository repositories.UserRepository, taskRepository repositories.TaskRepository) {
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

	user, err := userRepository.FindByID(ctx, claim.(*models.Claims).Subject)
	if err != nil {
		c.JSON(404, view.ErrorView("User not found"))
		return
	}

	taskId := c.Param("id")
	if taskId == "" {
		c.JSON(http.StatusBadRequest, view.ErrorView("Task ID is required"))
		return
	}

	task, err := taskRepository.FindByID(ctx, taskId)
	if err != nil {
		c.JSON(404, view.ErrorView("Task with id "+taskId+" not found"))
		return
	}

	if task.UserID != user.Id {
		c.JSON(http.StatusForbidden, view.ErrorView("You are not allowed to update this task"))
		return
	}

	task.Name = req.Name
	task.Description = req.Description
	task.StartsAt = req.StartsAt
	task.EndsAt = req.EndsAt

	if err := taskRepository.UpdateByID(ctx, task); err != nil {
		c.JSON(http.StatusInternalServerError, view.ErrorView("Failed to update task"))
		return
	}

	c.Status(http.StatusOK)
}

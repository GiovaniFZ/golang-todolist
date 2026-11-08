package controllers

import (
	"context"
	"encoding/json"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/view"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserTasks(c *gin.Context, db *gorm.DB) {
	ctx := context.Background()
	claim, exists := c.Get("userClaims")
	if !exists {
		c.JSON(404, view.ErrorView("userClaims not found"))
		return
	}

	user, err := gorm.G[models.User](db).Where("Id = ?", claim.(*models.Claims).Subject).First(ctx)
	if err != nil {
		c.JSON(404, view.ErrorView("User not found"))
		return
	}

	tasks, err := gorm.G[models.Task](db).Where("User_Id = ?", user.Id).Find(ctx)
	if err != nil {
		c.JSON(404, view.ErrorView("User tasks not found"))
		return
	}

	c.JSON(http.StatusOK, view.TaskView(tasks))
}

func CreateUserTask(c *gin.Context, db *gorm.DB) {
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

	user, err := gorm.G[models.User](db).Where("Id = ?", claim.(*models.Claims).Subject).First(ctx)
	if err != nil {
		c.JSON(404, view.ErrorView("User not found"))
		return
	}

	gorm.G[models.Task](db).Create(ctx, &models.Task{
		ID:          uuid.NewString(),
		UserId:      user.Id,
		Name:        req.Name,
		Description: req.Description,
		StartsAt:    req.StartsAt,
		EndsAt:      req.EndsAt,
	})

	c.Status(http.StatusCreated)
}

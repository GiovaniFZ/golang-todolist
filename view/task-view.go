package view

import (
	models "github/giovanifz/todolist-go/model"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskResponse struct {
	ID          string
	Description string
	Name        string
	StartsAt    time.Time
	EndsAt      time.Time
}

func TaskView(tasks []models.Task) gin.H {
	response := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		response = append(response, TaskResponse{
			ID:          task.ID,
			Description: task.Description,
			Name:        task.Name,
			StartsAt:    task.StartsAt,
			EndsAt:      task.EndsAt,
		})
	}
	return gin.H{
		"data": response,
	}
}

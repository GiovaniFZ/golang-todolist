package view

import (
	"fmt"
	models "github/giovanifz/todolist-go/model"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskResponse struct {
	Description string
	Name        string
	StartsAt    time.Time
	EndsAt      time.Time
}

func TaskView(tasks []models.Task) gin.H {
	response := make([]TaskResponse, 0, len(tasks))
	fmt.Println(tasks)
	for _, task := range tasks {
		response = append(response, TaskResponse{
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

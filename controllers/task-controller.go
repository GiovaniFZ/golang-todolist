package controllers

import (
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/view"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var mockTasks = []models.Task{
		{
			ID:          "1",
			Name:        "Homework",
			Description: "Math, Science and Chemistry",
			StartsAt:    time.Date(2027, time.November, 1, 0, 0, 0, 0, time.UTC),
			EndsAt:      time.Date(2027, time.November, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "2",
			Name:        "Do the dishes",
			Description: "Clean some things",
			StartsAt:    time.Date(2027, time.April, 10, 0, 0, 0, 0, time.UTC),
			EndsAt:      time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "3",
			Name:        "Go to school",
			Description: "Studies",
			StartsAt:    time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
			EndsAt:      time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	c.IndentedJSON(http.StatusOK, view.TaskView(mockTasks))
}

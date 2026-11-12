package mocks

import (
	"time"

	models "github/giovanifz/todolist-go/model"
)

var MockTask = models.Task{
	ID:          "1",
	Name:        "Homework",
	Description: "Math, Science and Chemistry",
	StartsAt:    time.Date(2027, time.November, 1, 0, 0, 0, 0, time.UTC),
	EndsAt:      time.Date(2027, time.November, 3, 0, 0, 0, 0, time.UTC),
}

package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      time.Time `json:"endsAt"`
}

var tasks = []task{
	{
		ID: "1", 
		Name: "Homework", 
		Description: "Math, Science and Chemistry", 
		StartsAt: time.Date(2027, time.November, 1, 0, 0, 0, 0, time.UTC), 
		EndsAt: time.Date(2027, time.November, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		ID: "2", 
		Name: "Do the dishes", 
		Description: "Clean some things", 
		StartsAt: time.Date(2027, time.April, 10, 0, 0, 0, 0, time.UTC),
		EndsAt: time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		ID: "3", 
		Name: "Go to school", 
		Description: "Studies", 
		StartsAt: time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
		EndsAt: time.Date(2027, time.November, 10, 0, 0, 0, 0, time.UTC),
	},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}

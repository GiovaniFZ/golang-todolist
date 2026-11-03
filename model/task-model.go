package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      time.Time `json:"endsAt"`
}

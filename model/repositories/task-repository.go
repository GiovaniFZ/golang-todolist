package repositories

import (
	"context"
	models "github/giovanifz/todolist-go/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	FindByUserID(ctx context.Context, userId string) ([]models.Task, error)
}

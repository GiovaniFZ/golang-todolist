package gorm

import (
	"context"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/model/repositories"

	"gorm.io/gorm"
)

type gormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &gormTaskRepository{db: db}
}

func (r *gormTaskRepository) Create(ctx context.Context, task *models.Task) error {
	gorm.G[models.Task](r.db).Create(ctx, task)
	if err := r.db.WithContext(ctx).Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormTaskRepository) FindByUserID(ctx context.Context, userId string) ([]models.Task, error) {
	var tasks []models.Task
	tasks, err := gorm.G[models.Task](r.db).Where("User_Id = ?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

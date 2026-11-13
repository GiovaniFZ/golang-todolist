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
	if err := gorm.G[models.Task](r.db).Create(ctx, task); err != nil {
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

func (r *gormTaskRepository) FindByID(ctx context.Context, id string) (*models.Task, error) {
	task, err := gorm.G[models.Task](r.db).Where("Id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *gormTaskRepository) UpdateByID(ctx context.Context, task *models.Task) error {
	_, err := gorm.G[models.Task](r.db).
		Where("id = ?", task.ID).
		Select("Name", "Description", "StartsAt", "EndsAt").
		Updates(ctx, *task)
	return err
}

func (r *gormTaskRepository) Delete(ctx context.Context, id string) error {
	_, err := gorm.G[models.Task](r.db).Where("Id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

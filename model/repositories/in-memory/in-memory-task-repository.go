package in_memory

import (
	"context"
	"errors"
	"sync"

	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/model/repositories"
)

var (
	errNilTask       = errors.New("task cannot be nil")
	errDuplicateTask = errors.New("task already exists")
)

type inMemoryTaskRepository struct {
	mu      sync.RWMutex
	tasks   map[string]models.Task
	taskIDs []string
}

func NewInMemoryTaskRepository() repositories.TaskRepository {
	return &inMemoryTaskRepository{
		tasks: make(map[string]models.Task),
	}
}

func (r *inMemoryTaskRepository) Create(ctx context.Context, task *models.Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if task == nil {
		return errNilTask
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errDuplicateTask
	}

	r.tasks[task.ID] = *task
	r.taskIDs = append(r.taskIDs, task.ID)
	return nil
}

func (r *inMemoryTaskRepository) FindByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]models.Task, 0)
	for _, taskID := range r.taskIDs {
		task := r.tasks[taskID]
		if task.UserId == userID {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

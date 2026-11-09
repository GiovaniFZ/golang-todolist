package in_memory

import (
	"context"
	"errors"
	"sync"

	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/model/repositories"
)

var (
	errNilUser       = errors.New("user cannot be nil")
	errDuplicateUser = errors.New("user already exists")
)

type inMemoryUserRepository struct {
	mu      sync.RWMutex
	users   map[string]models.User
	userIDs []string
}

func NewInMemoryUserRepository() repositories.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]models.User),
	}
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *models.User) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if user == nil {
		return errNilUser
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Id]; exists {
		return errDuplicateUser
	}

	r.users[user.Id] = *user
	r.userIDs = append(r.userIDs, user.Id)
	return nil
}

func (r *inMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	if err := ctx.Err(); err != nil {
		return &models.User{}, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, userID := range r.userIDs {
		user := r.users[userID]
		if user.Email == email {
			return &user, nil
		}
	}
	return &models.User{}, nil
}

func (r *inMemoryUserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	if err := ctx.Err(); err != nil {
		return &models.User{}, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return &models.User{}, nil
	}
	return &user, nil
}

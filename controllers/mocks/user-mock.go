package mocks

import models "github/giovanifz/todolist-go/model"

var MockUser = models.User{
	Name:  "John Doe",
	Email: "john.doe@example.com",
	Role:  "user",
	Tasks: []models.Task{},
}

var MockRegisterRequest = models.RegisterRequest{
	Name:     "John Doe",
	Email:    "john.doe@example.com",
	Password: "password123",
}

var MockLoginRequest = models.LoginRequest{
	Email:    "john.doe@example.com",
	Password: "password123",
}

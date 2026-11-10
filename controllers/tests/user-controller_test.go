package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github/giovanifz/todolist-go/controllers"
	"github/giovanifz/todolist-go/controllers/mocks"
	models "github/giovanifz/todolist-go/model"
	in_memory "github/giovanifz/todolist-go/model/repositories/in-memory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	userRepository := in_memory.NewInMemoryUserRepository()

	body, err := json.Marshal(mocks.MockRegisterRequest)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPost,
		"/sign-up",
		bytes.NewReader(body),
	)

	controllers.SignUp(c, userRepository)
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	savedUser, err := userRepository.FindByEmail(
		context.Background(),
		mocks.MockRegisterRequest.Email,
	)
	if err != nil {
		t.Fatal(err)
	}

	if savedUser.Name != mocks.MockRegisterRequest.Name {
		t.Errorf("Expected name %q, got %q",
			mocks.MockRegisterRequest.Name, savedUser.Name)
	}

}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	userRepository := in_memory.NewInMemoryUserRepository()
	password_hash := controllers.HashPassword(mocks.MockLoginRequest.Password)

	user := models.User{
		Name:          mocks.MockRegisterRequest.Name,
		Email:         mocks.MockRegisterRequest.Email,
		Role:          "user",
		Tasks:         []models.Task{},
		Password_Hash: password_hash,
	}

	userRepository.Create(
		context.Background(),
		&user,
	)

	body, err := json.Marshal(mocks.MockLoginRequest)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewReader(body),
	)

	controllers.Login(c, userRepository)
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var responseBodyJSON struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &responseBodyJSON)
	if err != nil {
		t.Fatal(err)
	}

	if responseBodyJSON.AccessToken == "" {
		t.Error("Expected to get an access token but got an empty access token")
	}
}

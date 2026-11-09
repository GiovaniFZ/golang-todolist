package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github/giovanifz/todolist-go/controllers"
	"github/giovanifz/todolist-go/controllers/mocks"
	in_memory "github/giovanifz/todolist-go/model/repositories/in-memory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	t.Setenv("JWT_SECRET", "mock-jwt-secret")

	userRepository := in_memory.NewInMemoryUserRepository()

	body, err := json.Marshal(mocks.MockRegisterRequest)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		http.MethodPost,
		"/users",
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

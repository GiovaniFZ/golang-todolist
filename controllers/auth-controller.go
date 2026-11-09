package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/model/repositories"
	"github/giovanifz/todolist-go/view"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenIssuer    = "auth-api"
	accessAudience = "auth-api:access"
	expirationTime = 15 * time.Minute
)

func GenerateAccessToken(id string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	now := time.Now()

	claims := models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tokenIssuer,
			Audience:  jwt.ClaimStrings{accessAudience},
			Subject:   fmt.Sprintf("%s", id),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expirationTime)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic("Could not hash password")
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context, userRepository repositories.UserRepository) {
	var req models.LoginRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView("Invalid body request"))
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, view.ErrorView("Email and password were not provided"))
		return
	}

	ctx := context.Background()
	user, err := userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView("User not found"))
		return
	}

	match := CheckPasswordHash(req.Password, user.Password_Hash)

	if !match {
		c.JSON(http.StatusBadRequest, view.ErrorView("Password invalid"))
		return
	}

	accessToken, err := GenerateAccessToken(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView(err.Error()))
		return
	}

	c.JSON(http.StatusOK, view.AuthSuccessView(accessToken, expirationTime))
}

func SignUp(c *gin.Context, userRepository repositories.UserRepository) {
	var req models.RegisterRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView("Invalid body request"))
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, view.ErrorView("Email and password were not provided"))
		return
	}

	ctx := context.Background()
	user, err := userRepository.FindByEmail(ctx, req.Email)

	if user != nil && user.Email == req.Email {
		c.JSON(http.StatusBadRequest, view.ErrorView("User already exists"))
		return
	}

	newUserId := uuid.NewString()
	passwordHash := HashPassword(req.Password)

	var newUser = models.User{
		Id:            newUserId,
		Name:          req.Name,
		Email:         req.Email,
		Role:          models.Member,
		Password_Hash: passwordHash,
	}

	userRepository.Create(ctx, &newUser)

	accessToken, err := GenerateAccessToken(newUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.ErrorView(err.Error()))
		return
	}
	c.JSON(http.StatusOK, view.AuthSuccessView(accessToken, expirationTime))
}

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/view"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func Login(c *gin.Context, db *gorm.DB) {
	var req models.LoginRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Invalid body request"))
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Email and password were not provided"))
		return
	}

	ctx := context.Background()
	user, err := gorm.G[models.User](db).Where("Email = ?", req.Email).First(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("User not found"))
		return
	}

	match := CheckPasswordHash(req.Password, user.Password_Hash)

	if !match {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Password invalid"))
		return
	}

	accessToken, err := GenerateAccessToken(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView(err.Error()))
		return
	}

	c.JSON(http.StatusOK, view.AuthSuccessView(accessToken, expirationTime))
}

func SignUp(c *gin.Context, db *gorm.DB) {
	var req models.RegisterRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Invalid body request"))
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Email and password were not provided"))
		return
	}

	ctx := context.Background()
	user, err := gorm.G[models.User](db).Where("Email = ?", req.Email).First(ctx)

	if user.Email == req.Email {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("User already exists"))
		return
	}

	gorm.G[models.User](db).Create(ctx, &models.User{
		Id:            uuid.NewString(),
		Name:          req.Name,
		Email:         req.Email,
		Role:          models.Member,
		Password_Hash: HashPassword(req.Password),
	})

	accessToken, err := GenerateAccessToken(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView(err.Error()))
		return
	}
	c.JSON(http.StatusOK, view.AuthSuccessView(accessToken, expirationTime))
}

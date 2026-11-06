package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	models "github/giovanifz/todolist-go/model"
	"github/giovanifz/todolist-go/view"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

const (
	tokenIssuer    = "auth-api"
	accessAudience = "auth-api:access"
	expirationTime = 15 * time.Minute
)

func getSecret() ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		return nil, err
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT Secret missing")
	}
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil || len(key) < 32 {
		return nil, errors.New("JWT_SECRET should be Base64 with 32 bytes")
	}
	return key, nil
}

func GenerateAccessToken(id string) (string, error) {
	secret, err := getSecret()
	fmt.Println(secret)
	if err != nil {
		return "", err
	}

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

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("Error: %w", err)
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (*models.Claims, error) {
	secret, err := getSecret()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}, jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(tokenIssuer),
		jwt.WithAudience(accessAudience))

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("Expired token")
		}
		return nil, errors.New("Expired token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}

func Login(c *gin.Context) {
	var req models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Invalid body request"))
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, view.AuthErrorView("Email and password were not provided"))
		return
	}

	accessToken, err := GenerateAccessToken(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.AuthErrorView(err.Error()))
		return
	}

	c.JSON(http.StatusOK, view.AuthSuccessView(accessToken, expirationTime))
}

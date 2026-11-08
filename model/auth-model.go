package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type roles string

const (
	Admin  roles = "admin"
	Member roles = "member"
)

type User struct {
	Id            string `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"not null" json:"name"`
	Email         string `gorm:"unique;not null" json:"email"`
	Role          roles  `gorm:"not null" json:"role"`
	Password_Hash string `gorm:"not null"`
	Tasks         []Task `gorm:"foreignKey:UserId"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
}

type ErrorResponse struct {
	Error string `json:"erro"`
	Code  int    `json:"code"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

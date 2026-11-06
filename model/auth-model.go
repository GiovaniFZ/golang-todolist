package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
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

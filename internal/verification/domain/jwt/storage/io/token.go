package io

import "github.com/golang-jwt/jwt/v4"

type GenerateTokenInput struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Salt     string `json:"salt"`
	jwt.RegisteredClaims
}

type Token struct {
	UserID      string
	Username    string
	Token       string
	TokenExpiry int64
}

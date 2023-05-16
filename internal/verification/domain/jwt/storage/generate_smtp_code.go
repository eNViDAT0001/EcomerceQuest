package storage

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtStorage) GenerateSmtpCodeVerification(ctx context.Context, code string, email string) (string, error) {

	claims := jwt.MapClaims{}
	claims["email"] = email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := code
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

package storage

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtStorage) VerifySmtpToken(ctx context.Context, tokenString string, signature string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signature), nil
	})
}

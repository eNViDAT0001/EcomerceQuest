package storage

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtStorage) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(wViper.GetString("TOKEN_SECRET")), nil
	})
}

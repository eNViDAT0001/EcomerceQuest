package storage

import (
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtStorage) GenerateSmtpCodeVerification(code string) (string, error) {

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, nil)
	secret := code
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

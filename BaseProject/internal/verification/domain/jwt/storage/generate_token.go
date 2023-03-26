package storage

import (
	"context"
	"time"

	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtStorage) GenerateToken(ctx context.Context, input io.GenerateTokenInput, expiresAt time.Time) (*io.Token, error) {

	// TODO: Check claims
	//claims := jwt.MapClaims{}
	//claims["salt"] = input.Salt
	//claims["username"] = input.Username
	//claims["time"] = time.Now()
	//claims["ExpiresAt"] = expiresAt.Unix()
	input.ExpiresAt = jwt.NewNumericDate(expiresAt)
	input.IssuedAt = jwt.NewNumericDate(time.Now())
	input.NotBefore = jwt.NewNumericDate(time.Now())

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, input)
	secret := wViper.GetString("TOKEN_SECRET")
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	data := &io.Token{
		Username:    input.Username,
		Token:       token,
		TokenExpiry: expiresAt.Unix(),
	}

	return data, nil
}

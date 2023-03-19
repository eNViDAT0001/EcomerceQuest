package usecase

import (
	"context"
	"math/rand"
	"time"
)

func (s *jwtUseCase) GenerateSmtpCode(ctx context.Context, userID uint) (token string, code string, err error) {
	user, err := s.userSto.GetUserDetailByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	rand.Seed(time.Now().UnixNano())
	code = string(rune(rand.Intn(999999-100000) + 100000))
	signature := code + *user.Salt
	token, err = s.tokenSto.GenerateSmtpCodeVerification(signature)
	if err != nil {
		return "", "", err
	}
	return token, code, nil
}

package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (s *jwtUseCase) GenerateSmtpCode(ctx context.Context, email string) (token string, code string, err error) {
	user, err := s.userSto.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	rand.Seed(time.Now().UnixNano())
	code = strconv.Itoa(rand.Intn(999999-100000) + 100000)
	signature := fmt.Sprintf("%s %s", *user.Salt, code)
	token, err = s.tokenSto.GenerateSmtpCodeVerification(nil, signature, *user.Email)
	if err != nil {
		return "", "", err
	}
	return token, code, nil
}

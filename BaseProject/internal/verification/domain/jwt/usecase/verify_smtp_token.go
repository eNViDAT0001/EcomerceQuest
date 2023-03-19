package usecase

import (
	"context"
)

func (s *jwtUseCase) VerifySmtpToken(ctx context.Context, userID uint, accessToken string, code string) (bool, error) {
	user, err := s.userSto.GetUserDetailByID(ctx, userID)
	if err != nil {
		return false, err
	}

	signature := *user.Salt + code
	token, err := s.tokenSto.VerifySmtpToken(accessToken, signature)
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}

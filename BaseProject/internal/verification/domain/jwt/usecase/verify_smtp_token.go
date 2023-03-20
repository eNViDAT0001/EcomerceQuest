package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtUseCase) VerifySmtpToken(ctx context.Context, userID uint, accessToken string, code string) (*jwt.Token, error) {
	user, err := s.userSto.GetUserDetailByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	signature := fmt.Sprintf("%s %s", *user.Salt, code)
	token, err := s.tokenSto.VerifySmtpToken(accessToken, signature)
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	email, _ := claims["email"].(string)
	if email != *user.Email {
		return nil, fmt.Errorf("invalid email")
	}

	return token, nil
}

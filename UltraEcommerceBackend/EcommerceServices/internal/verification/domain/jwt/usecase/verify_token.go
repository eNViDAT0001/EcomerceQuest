package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
	"github.com/golang-jwt/jwt/v4"
)

func (s *jwtUseCase) VerifyToken(ctx context.Context, accessToken string) (*jwt.Token, error) {
	token, err := s.tokenSto.VerifyToken(accessToken)
	if err != nil && !token.Valid {
		return nil, request.NewUnauthorizedError("token", accessToken, "Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, request.NewUnauthorizedError("token", accessToken, "Invalid Token")
	}
	userIdentify := io.UserIdentify{
		Username: claims["username"].(string),
	}

	_, err = s.userSto.GetUserWithIdentify(ctx, userIdentify)
	if err != nil {
		return nil, request.NewUnauthorizedError("refresh_token", token, "Invalid Refresh Token")
	}

	return token, nil
}

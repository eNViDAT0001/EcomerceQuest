package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase/io"
)

func (s *jwtUseCase) GenerateToken(ctx context.Context, input io.GenerateTokenInput) (*ioUC.JwtToken, error) {
	days, _ := ioUC.GetDate()
	minutes, _ := ioUC.GetMinute()

	accessToken, err := s.tokenSto.GenerateToken(nil, input, minutes)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenSto.GenerateToken(nil, input, days)
	if err != nil {
		return nil, err
	}

	output := ioUC.JwtToken{
		AccessToken:        accessToken.Token,
		AccessTokenExpiry:  accessToken.TokenExpiry,
		RefreshToken:       refreshToken.Token,
		RefreshTokenExpiry: refreshToken.TokenExpiry,
	}

	return &output, nil
}

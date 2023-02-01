package usecase

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase/io"
)

func (s *jwtUseCase) GenerateToken(input io.GenerateTokenInput) (*ioUC.JwtToken, error) {
	days, _ := ioUC.GetDate()
	minutes, _ := ioUC.GetMinute()

	accessToken, err := s.tokenSto.GenerateToken(input, minutes)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenSto.GenerateToken(input, days)
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

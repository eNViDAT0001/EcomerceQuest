package jwt

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase/io"
	"github.com/golang-jwt/jwt/v4"
)

type UseCase interface {
	GenerateToken(input io.GenerateTokenInput) (*ioUC.JwtToken, error)
	RefreshToken(ctx context.Context, refreshToken string) (*ioUC.JwtToken, error)
	GenerateSmtpCode(ctx context.Context, userID uint) (token string, code string, err error)
	VerifyToken(ctx context.Context, accessToken string) (*jwt.Token, error)
	VerifySmtpToken(ctx context.Context, userID uint, accessToken string, code string) (bool, error)
}

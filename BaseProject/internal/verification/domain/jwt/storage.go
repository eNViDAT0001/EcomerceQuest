package jwt

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Storage interface {
	GenerateToken(ctx context.Context, input io.GenerateTokenInput, expiresAt time.Time) (*io.Token, error)
	GenerateSmtpCodeVerification(ctx context.Context, code string, email string) (string, error)
	VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, error)
	VerifySmtpToken(ctx context.Context, tokenString string, signature string) (*jwt.Token, error)
}

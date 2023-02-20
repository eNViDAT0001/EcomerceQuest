package app_accession

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/verification/domain/jwt/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/verification/domain/jwt/usecase/io"
)

type UseCase interface {
	Login(ctx context.Context, username string, password string) (*entities.User, error)
	LoginByGoogle(ctx context.Context, client io.Client) (*ioUC.JwtToken, error)
	LoginByFacebook(ctx context.Context, client io.Client) (*ioUC.JwtToken, error)
}

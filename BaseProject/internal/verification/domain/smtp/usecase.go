package smtp

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
)

type UseCase interface {
	SendEmail(ctx context.Context, input io.EmailForm) error
	CreateEmail(ctx context.Context, email ioSto.CreateEmail) (id uint, err error)
}

package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
)

type Storage interface {
	CreateEmail(ctx context.Context, email io.CreateEmail) (id uint, err error)
}

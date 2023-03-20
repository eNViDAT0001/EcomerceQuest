package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
)

type UseCase interface {
	SendEmail(ctx context.Context, input io.EmailForm) error
}

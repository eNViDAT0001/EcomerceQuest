package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
)

func (s *smtpUseCase) CreateEmail(ctx context.Context, email io.CreateEmail) (id uint, err error) {
	return s.smtpStorage.CreateEmail(ctx, email)

}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"gorm.io/gorm"
)

func (s *smtpUseCase) ListEmails(ctx context.Context, filter paging.ParamsInput) (emails []entities.Email, total int64, err error) {
	total, err = s.smtpStorage.CountListEmail(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	emails, err = s.smtpStorage.ListEmail(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return emails, total, err
}

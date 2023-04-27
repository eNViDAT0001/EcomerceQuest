package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

func (s smtpStorage) CountListEmail(ctx context.Context, email io.CreateEmail) (id uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.Email{}.TableName()).Create(&email).Error
	if err != nil {
		return 0, err
	}
	return email.ID, nil
}

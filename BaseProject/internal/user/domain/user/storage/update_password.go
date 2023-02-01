package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userStorage) UpdatePassword(ctx context.Context, userID uint32, password string) error {
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	hashedPassword, err := entities.User{}.HashPassword(password)

	if err != nil {
		return err
	}

	err = db.Table(tableName).
		Where("id = ?", userID).
		Update("password", hashedPassword).
		Where("deleted_at IS NULL").
		Error

	if err != nil {
		return err
	}

	return nil
}

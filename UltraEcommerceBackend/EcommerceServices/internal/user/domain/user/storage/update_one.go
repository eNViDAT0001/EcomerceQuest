package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userStorage) UpdateUser(ctx context.Context, userID uint32, input *io.UpdateUserInput) error {
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Where("id = ?", userID).
		Where("deleted_at IS NULL").
		Updates(&input).Error

	if err != nil {

		return err
	}

	return nil
}

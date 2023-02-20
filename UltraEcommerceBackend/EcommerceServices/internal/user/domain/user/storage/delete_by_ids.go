package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userStorage) DeleteUserByIDs(ctx context.Context, IDs []uint) error {
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("id IN ?", IDs).Delete(&entities.User{}).Error

	if err != nil {
		return err
	}

	return nil
}

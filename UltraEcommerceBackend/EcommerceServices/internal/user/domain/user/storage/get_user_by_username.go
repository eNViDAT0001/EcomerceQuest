package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userStorage) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	result := entities.User{}
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("username = ?", username).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userStorage) GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error) {
	result := entities.User{}
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("id = ?", ID).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

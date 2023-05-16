package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userStorage) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	result := entities.User{}
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("email = ?", email).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

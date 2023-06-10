package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userStorage) GetListByProviderID(ctx context.Context, IDs []uint) ([]entities.User, error) {
	result := make([]entities.User, 0)
	db := wrap_gorm.GetDB()

	err := db.Model(entities.User{}).
		Joins("LEFT JOIN Provider ON Provider.user_id = User.id").
		Where("Provider.id IN ?", IDs).
		Find(&result).Error

	return result, err
}
func (u userStorage) GetDetailByProviderID(ctx context.Context, id uint) (entities.User, error) {
	result := entities.User{}
	db := wrap_gorm.GetDB()

	err := db.Model(entities.User{}).
		Joins("JOIN Provider ON Provider.user_id = User.id").
		Where("Provider.id = ?", id).
		First(&result).
		Group("User.id").
		Error

	return result, err
}

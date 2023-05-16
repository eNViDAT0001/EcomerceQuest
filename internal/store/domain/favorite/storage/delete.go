package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (f favoriteStorage) DeleteFavorite(ctx context.Context, userID uint, providerID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Where("user_id = ?", userID).Where("provider_id = ?", providerID).Delete(&entities.Provider{}).Error
	if err != nil {
		return err
	}
	return nil
}

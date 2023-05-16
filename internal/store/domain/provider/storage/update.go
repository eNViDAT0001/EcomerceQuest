package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) UpdateProvider(ctx context.Context, providerID uint, input io.ProviderUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Provider{}).Where("id = ?", providerID).Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

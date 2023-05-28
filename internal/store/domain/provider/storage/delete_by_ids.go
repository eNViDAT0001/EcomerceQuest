package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) DeleteProviderByIDs(ctx context.Context, providerID []uint) error {
	tableName := entities.Provider{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).
		Where("id IN ?", providerID).
		Delete(&entities.Provider{}).Error
	if err != nil {
		return err
	}

	err = db.Table(entities2.Product{}.TableName()).
		Where("provider_id IN ?", providerID).
		Delete(&entities2.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}

package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	entities3 "github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (s providerStorage) DeleteProviderByIDs(ctx context.Context, providerID []uint) error {
	tableName := entities.Provider{}.TableName()
	database := wrap_gorm.GetDB()

	err := database.Transaction(func(db *gorm.DB) error {
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

		err = db.Table(entities3.Order{}.TableName()).
			Where("provider_id IN ?", providerID).
			Where("status != ?", entities3.DeliveredOrder).
			Where("status != ?", entities3.CancelOrder).
			Update("status", entities3.CancelOrder).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (b couponStorage) UpdateBanner(ctx context.Context, bannerID uint, input io.BannerUpdateForm, productIDsIN []uint, productIDsOUT []uint) error {
	database := wrap_gorm.GetDB()

	err := database.Transaction(func(db *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		err := db.Model(entities.Banner{}).Where("id = ?", bannerID).Updates(&input).Error
		if err != nil {
			return err
		}

		productINStorage := make([]entities.BannerDetail, 0)
		for _, id := range productIDsIN {
			bannerDetail := entities.BannerDetail{
				BannerID:  bannerID,
				ProductID: id,
			}
			productINStorage = append(productINStorage, bannerDetail)
		}
		if len(productINStorage) > 0 {
			err = db.Table(entities.BannerDetail{}.TableName()).Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(&productINStorage).Error
			if err != nil {
				db.Rollback()
				return err
			}
		}

		if len(productIDsOUT) > 0 {
			err = db.Table(entities.BannerDetail{}.TableName()).
				Where("product_id IN ?", productIDsOUT).
				Delete(&entities.BannerDetail{}).
				Error

			if err != nil {
				db.Rollback()
				return err
			}
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

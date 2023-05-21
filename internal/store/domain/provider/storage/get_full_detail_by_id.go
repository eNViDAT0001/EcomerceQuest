package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
	entities3 "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	entities2 "github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) GetProviderFullDetailByID(ctx context.Context, id uint) (io.ProviderFullDetail, error) {
	var result io.ProviderFullDetail
	db := wrap_gorm.GetDB()

	err := db.Model(entities2.Provider{}).Where("id = ?", id).First(&result.Provider).Error
	if err != nil {
		return result, err
	}
	err = db.Model(entities.Order{}).
		Where("provider_id = ?", id).
		Group("provider_id").
		Where("status = ?", entities.DeliveredOrder).
		Count(&result.TotalOrders).Error
	if err != nil {
		return result, err
	}
	err = db.Model(entities.Order{}).
		Select("SUM(total)").
		Where("provider_id = ?", id).
		Where("status = ?", entities.DeliveredOrder).
		Group("provider_id").
		Scan(&result.Revenue).Error
	if err != nil {
		return result, err
	}
	err = db.Model(entities3.Comment{}).
		Select("ROUND(AVG(`Comment`.rating),0)").
		Joins("JOIN Product ON Product.id = Comment.product_id").
		Joins("JOIN Provider ON Provider.id = Product.provider_id").
		Where("Provider.id = ?", id).
		Group("Provider.id").
		Scan(&result.AverageComment).Error
	if err != nil {
		return result, err
	}
	err = db.Model(entities3.Comment{}).
		Joins("JOIN Product ON Product.id = Comment.product_id").
		Joins("JOIN Provider ON Provider.id = Product.provider_id").
		Where("Provider.id = ?", id).
		Group("Provider.id").
		Count(&result.TotalComment).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

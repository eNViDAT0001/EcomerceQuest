package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s orderStorage) ListByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) ([]entities.Order, error) {
	result := make([]entities.Order, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).Where("provider_id = ?", providerID)
	paging_query.SetPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (s providerStorage) ListProviderByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) ([]entities.Provider, error) {
	result := make([]entities.Provider, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Provider{}).Where("user_id = ?", userID)
	paging_query.SetPagingQuery(&filter, entities.Provider{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

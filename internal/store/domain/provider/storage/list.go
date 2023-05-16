package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) ListProvider(ctx context.Context, filter paging.ParamsInput) ([]entities.Provider, error) {
	result := make([]entities.Provider, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Provider{})
	paging_query.SetPagingQuery(&filter, entities.Provider{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

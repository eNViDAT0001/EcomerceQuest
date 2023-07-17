package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) ListProviderQuantity(ctx context.Context, filter paging.ParamsInput) ([]io.ProviderQuantity, error) {
	result := make([]io.ProviderQuantity, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Provider{}).
		Select("(created_at AS date) as date, COUNT(id) as quantity").
		Group("CAST(created_at AS date)")
	paging_query.SetPagingQuery(&filter, entities.Provider{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

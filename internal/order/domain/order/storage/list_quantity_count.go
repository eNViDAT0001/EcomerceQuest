package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s *orderStorage) ListQuantityCount(ctx context.Context, input paging.ParamsInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).
		Group("CAST(created_at AS date)")
	paging_query.SetCountListPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

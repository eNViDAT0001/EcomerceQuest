package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

func (s *orderStorage) ListQuantity(ctx context.Context, input paging.ParamsInput) ([]io.OrderReportQuantity, error) {
	result := make([]io.OrderReportQuantity, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).
		Select("CAST(created_at AS date) as date, SUM(quantity) as quantity, SUM(total) as total").
		Group("CAST(created_at AS date)")
	paging_query.SetPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

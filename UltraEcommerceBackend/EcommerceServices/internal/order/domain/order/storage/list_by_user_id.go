package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/entities"
)

func (s *orderStorage) ListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]entities.Order, error) {
	result := make([]entities.Order, 0)
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Order{}).Where("user_id = ?", userID)
	paging_query.SetPagingQuery(&input, entities.Order{}.TableName(), query)
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

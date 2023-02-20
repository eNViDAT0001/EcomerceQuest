package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/entities"
)

func (c cartStorage) CountListCartByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (total int64, err error) {
	var count int64
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartDetail{}.TableName()).Where("user_id = ?", userID)

	paging_query.SetCountListPagingQuery(&filter, entities.CartDetail{}.TableName(), query)

	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, err
}

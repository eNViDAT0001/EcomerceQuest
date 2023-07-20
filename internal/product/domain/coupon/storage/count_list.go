package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) CountListCoupon(ctx context.Context, filter paging.ParamsInput, couponID uint) (total int64, err error) {
	var count int64

	db := wrap_gorm.GetDB()

	query := db.Model(entities.Coupon{})

	paging_query.SetCountListPagingQuery(&filter, entities.Coupon{}.TableName(), query)

	if couponID > 0 {
		query = query.Where("id = ?", couponID)
	}

	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

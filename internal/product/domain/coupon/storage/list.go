package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (b couponStorage) ListCoupon(ctx context.Context, filter paging.ParamsInput) ([]entities.Coupon, error) {
	result := make([]entities.Coupon, 0)

	db := wrap_gorm.GetDB()

	query := db.Model(entities.Coupon{})

	paging_query.SetPagingQuery(&filter, entities.Coupon{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) ListCountRecommenderProductsPreview(ctx context.Context, input io.ListRecommendedProductInput) (total int, err error) {
	var result int64
	db := wrap_gorm.GetDB()

	query := db.Table(entities.ProductPreview{}.TableName())
	paging_query.SetCountListPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), query)
	err = query.Count(&result).Error
	if err != nil {
		return 0, err
	}
	total = int(result)

	return total, nil
}

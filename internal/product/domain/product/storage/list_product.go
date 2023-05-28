package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) ListProduct(ctx context.Context, input io.ListProductInput) ([]io.ProductWithQuantities, error) {
	result := make([]io.ProductWithQuantities, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Product{}).
		Select("Product.*, IF(COUNT(ProductOption.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT('name', ProductOption.name, 'quantity', ProductOption.quantity))) AS options").
		Joins("LEFT JOIN ProductOption on Product.id = ProductOption.product_id").
		Joins("JOIN Provider ON Product.provider_id = Provider.id AND Provider.deleted_at IS NULL").
		Where("ProductOption.deleted_at IS NULL").
		Group("Product.id")

	paging_query.SetPagingQuery(&input.Paging, entities.Product{}.TableName(), query)

	if len(input.ProductIDs) > 0 {
		query = query.Where("Product.id IN ?", input.ProductIDs)
	}
	if len(input.CategoryIDs) > 0 {
		query = query.Where("Product.category_id IN ?", input.CategoryIDs)
	}

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

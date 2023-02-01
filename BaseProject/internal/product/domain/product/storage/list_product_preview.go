package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (s productStorage) ListProductsPreview(ctx context.Context, input io.ListProductInput) ([]io.ProductPreviewItem, error) {
	result := make([]io.ProductPreviewItem, 0)
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Product{}.TableName()).
		Select(
			"Product.id, " +
				"Product.provider_id, " +
				"Product.category_id, " +
				"Product.user_id, " +
				"Product.name, " +
				"Product.price, " +
				"Product.discount, " +
				"IF(COUNT(ProductMedia.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'publicID', ProductMedia.public_id, 'mediaPath', ProductMedia.media_path, 'type', ProductMedia.media_type))) AS media, " +
				"AVG(Comment.rating) AS rating").
		Joins("LEFT JOIN ProductMedia ON ProductMedia.product_id = Product.id").
		Joins("LEFT JOIN Comment ON Comment.product_id = Product.id").
		Where("Product.deleted_at IS NULL").
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

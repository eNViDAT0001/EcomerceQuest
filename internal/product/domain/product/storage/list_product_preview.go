package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
	"strconv"
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
				"Product.short_descriptions, " +
				"Product.height, " +
				"Product.weight, " +
				"Product.width, " +
				"Product.length, " +
				"IF(COUNT(ProductMedia.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'publicID', ProductMedia.public_id, 'mediaPath', ProductMedia.media_path, 'type', ProductMedia.media_type))) AS media, " +
				"ROUND(AVG(Comment.rating),0) AS rating").
		Joins("LEFT JOIN ProductMedia ON ProductMedia.product_id = Product.id").
		Joins("LEFT JOIN Comment ON Comment.product_id = Product.id").
		Where("Product.deleted_at IS NULL").
		Group("Product.id")

	err := DoDummyRatingFilter(input, query) // This is the DummyLine
	if err != nil {
		return nil, err
	}
	paging_query.SetPagingQuery(&input.Paging, entities.Product{}.TableName(), query)

	if len(input.ProductIDs) > 0 {
		query = query.Where("Product.id IN ?", input.ProductIDs)
	}
	if len(input.CategoryIDs) > 0 {
		query = query.Where("Product.category_id IN ?", input.CategoryIDs)
	}

	err = query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	RemoveDummyFilter(input) // This is the DummyLine
	return result, nil
}

var tempStorage *string

func DoDummyRatingFilter(input io.ListProductInput, query *gorm.DB) error {
	fields := input.Paging.Filter.GetFields()
	if fields == nil {
		return nil
	}

	val, ok := (*fields)["rating"]
	if !ok {
		return nil
	}
	tempStorage = &val
	rating, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	query = query.Having("ROUND(AVG(`Comment`.`rating`),0) = ?", rating)
	delete(*fields, "rating")
	return nil
}
func RemoveDummyFilter(input io.ListProductInput) {
	fields := input.Paging.Filter.GetFields()
	if tempStorage != nil && fields != nil {
		(*fields)["rating"] = *tempStorage
	}
}

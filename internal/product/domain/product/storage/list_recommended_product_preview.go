package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

// 1. Get all recommended products (rp)
// 2. Compare len(rp) with total products
// 3. If len(rp) < total products, get random products
// 4. If len(rp) > total products, left get recommended products
// 5. If len(rp) > total products and left < 0 get random products
func (s productStorage) ListRecommendedProductsPreview(ctx context.Context, input io.ListRecommendedProductInput) ([]entities.ProductPreview, error) {
	result := make([]entities.ProductPreview, 0)
	db := wrap_gorm.GetDB()
	if input.Paging.Marker < 1 {
		input.Paging.Marker = 1
	}
	if input.Paging.Limit < 1 {
		input.Paging.Limit = 10
	}

	var count int64
	countQuery := db.Table(entities.ProductPreview{}.TableName()).Where("id IN ?", input.RecommendedProductIDs)
	paging_query.SetCountListPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), countQuery)
	err := countQuery.Count(&count).Error
	if err != nil {
		return nil, err
	}
	var totalRecommendedProduct = int(count)

	// Nếu có marker trong khoảng recommend
	require := input.Paging.Marker * input.Paging.Limit
	if require <= totalRecommendedProduct {
		query := db.Table(entities.ProductPreview{}.TableName()).Where("id IN ?", input.RecommendedProductIDs)
		paging_query.SetPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), query)
		err := query.Scan(&result).Error
		if err != nil {
			return nil, err
		}

		return result, err
	}

	// Nếu marker ở giữa Recommended và Chưa đc recommend
	left := require - totalRecommendedProduct
	if left < input.Paging.Limit {
		query := db.Table(entities.ProductPreview{}.TableName()).Where("id IN ?", input.RecommendedProductIDs)
		paging_query.SetPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), query)
		err := query.Scan(&result).Error
		if err != nil {
			return nil, err
		}

		unRecommended := make([]entities.ProductPreview, 0)
		tempLimit := input.Paging.Limit
		input.Paging.Limit = input.Paging.Limit - len(result)
		tempMarker := input.Paging.Marker
		input.Paging.Marker = 0
		query = db.Table(entities.ProductPreview{}.TableName()).
			Where("id NOT IN ?", input.RecommendedProductIDs).
			Order("rating DESC")
		paging_query.SetPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), query)
		err = query.Scan(&unRecommended).Error
		if err != nil {
			return nil, err
		}
		input.Paging.Marker = tempMarker
		input.Paging.Limit = tempLimit
		return append(result, unRecommended...), err
	}

	// Nếu có marker tại khúc và Chưa đc recommend

	recommendPages := totalRecommendedProduct / input.Paging.Limit
	if totalRecommendedProduct%input.Paging.Limit != 0 {
		recommendPages++
	}

	tempMarker := input.Paging.Marker
	input.Paging.Marker = input.Paging.Marker - recommendPages

	query := db.Table(entities.ProductPreview{}.TableName()).
		Where("id NOT IN ?", input.RecommendedProductIDs).
		Order("rating DESC")
	paging_query.SetPagingQuery(&input.Paging, entities.ProductPreview{}.TableName(), query)
	query = query.Offset(DummyOffset(input.Paging.Marker, input.Paging.Limit, (require-totalRecommendedProduct)%input.Paging.Limit))
	err = query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	input.Paging.Marker = tempMarker
	return result, err
}
func DummyOffset(pageNum int, pageSize int, left int) int {
	if pageNum == 1 {
		return left
	}
	offset := (pageNum - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	return offset + left
}

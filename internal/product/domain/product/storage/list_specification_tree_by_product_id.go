package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"

	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"gorm.io/gorm"
)

func (s productStorage) GetSpecificationTreeByProductID(ctx context.Context, productID uint) ([]io.ProductSpecificationFullDetail, error) {

	result := make([]io.ProductSpecificationFullDetail, 0)
	db := wrap_gorm.GetDB()

	err := db.Model(entities.ProductSpecification{}).
		Select("ProductSpecification.id, ProductSpecification.properties, JSON_ARRAYAGG(JSON_OBJECT( 'id', ProductOption.id, 'name', ProductOption.name, 'price', ProductOption.price, 'quantity', ProductOption.quantity)) AS options").
		Joins("JOIN ProductOption ON ProductOption.specification_id = ProductSpecification.id").
		Joins("JOIN Product ON Product.id = ProductSpecification.product_id").
		Where("ProductSpecification.product_id", productID).
		Group("ProductSpecification.id").
		Scan(&result).Error

	if err != nil {
		return result, err
	}
	if len(result) < 1 {
		return result, gorm.ErrRecordNotFound
	}
	return result, err
}

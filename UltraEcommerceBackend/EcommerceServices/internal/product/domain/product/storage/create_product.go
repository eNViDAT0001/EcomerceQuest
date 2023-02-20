package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (s productStorage) CreateProduct(ctx context.Context, product io.ProductCreateForm) (productID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.Product{}.TableName()).
		Create(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (u *productUseCase) ListRecommendedProductsPreview(ctx context.Context, input ioSto.ListRecommendedProductInput) (items []entities.ProductPreview, total int64, err error) {
	count, err := u.productSto.ListCountRecommenderProductsPreview(ctx, input)
	total = int64(count)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	input.Count = count
	products, err := u.productSto.ListRecommendedProductsPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

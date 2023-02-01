package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

func (u *productUseCase) ListProduct(ctx context.Context, input ioSto.ListProductInput) (items []entities.Product, total int64, err error) {
	total, err = u.productSto.ListCountProduct(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err := u.productSto.ListProduct(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

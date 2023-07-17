package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"gorm.io/gorm"
)

func (u *productUseCase) ListProductQuantity(ctx context.Context, input ioSto.ListProductInput) (items []ioSto.ProductQuantity, total int64, err error) {
	total, err = u.productSto.ListCountProductQuantity(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err := u.productSto.ListProductQuantity(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

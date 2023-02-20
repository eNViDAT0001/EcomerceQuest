package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
)

func (u *productUseCase) CreateProductOptions(ctx context.Context, input []ioSto.ProductOptionCreateForm) error {
	return u.productSto.CreateProductOptions(ctx, input)
}

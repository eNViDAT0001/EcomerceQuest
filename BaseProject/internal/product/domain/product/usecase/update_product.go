package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
)

func (u *productUseCase) UpdateProduct(ctx context.Context, productID uint, product ioSto.ProductUpdateForm) error {
	return u.productSto.UpdateProduct(ctx, productID, product)
}

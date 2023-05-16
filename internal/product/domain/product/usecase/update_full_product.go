package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
)

func (u *productUseCase) UpdateFullProduct(ctx context.Context, productID uint, product ioSto.ProductFullUpdateForm) error {
	return u.productSto.UpdateFullProduct(ctx, productID, product)
}

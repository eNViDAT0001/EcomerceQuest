package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
)

func (u *productUseCase) UpdateProductSpecification(ctx context.Context, specID uint, specifications ioSto.ProductSpecificationUpdateForm) error {
	return u.productSto.UpdateProductSpecification(ctx, specID, specifications)
}

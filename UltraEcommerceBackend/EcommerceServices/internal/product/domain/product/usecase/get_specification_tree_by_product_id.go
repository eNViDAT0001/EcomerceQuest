package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
)

func (u *productUseCase) GetSpecificationTreeByProductID(ctx context.Context, productID uint) ([]ioSto.ProductSpecificationFullDetail, error) {
	return u.productSto.GetSpecificationTreeByProductID(ctx, productID)
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (u *productUseCase) ListProductMediaByProductID(ctx context.Context, productID uint) ([]entities.ProductMedia, error) {
	return u.productSto.ListProductMediaByProductID(ctx, productID)
}

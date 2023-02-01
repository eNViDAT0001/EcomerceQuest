package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (u *productUseCase) GetProductInfoByID(ctx context.Context, productID uint) (entities.Product, error) {
	return u.productSto.GetProductDetailByID(ctx, productID)
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (u *productUseCase) GetProductDescriptionsByProductID(ctx context.Context, productID uint) ([]entities.ProductDescriptions, error) {
	return u.productSto.GetProductDescriptionsByProductID(ctx, productID)
}

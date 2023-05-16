package usecase

import "context"

func (u *productUseCase) DeleteProductByID(ctx context.Context, ID uint) error {
	return u.productSto.DeleteProductByID(ctx, ID)
}

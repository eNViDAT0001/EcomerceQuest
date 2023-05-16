package usecase

import "context"

func (u *productUseCase) DeleteProductByIDs(ctx context.Context, IDs []uint) error {
	return u.productSto.DeleteProductByIDs(ctx, IDs)
}

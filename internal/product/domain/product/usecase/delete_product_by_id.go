package usecase

import "context"

func (u *productUseCase) DeleteProductByID(ctx context.Context, ID uint) error {
	return u.productSto.DeleteProductByID(ctx, ID)
}
func (u *productUseCase) DeleteElementByIDs(ctx context.Context, ID uint, descriptionsIDs []uint, mediaIDs []uint) error {
	return u.productSto.DeleteElementByIDs(ctx, ID, descriptionsIDs, mediaIDs)
}

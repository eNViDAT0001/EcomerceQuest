package usecase

import "context"

func (u *cartUseCase) DeleteCart(ctx context.Context, userID uint, cartID uint) error {
	return u.cartSto.DeleteCart(ctx, userID, cartID)
}

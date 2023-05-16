package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/entities"
)

func (u *cartUseCase) GetDetailByID(ctx context.Context, cartID uint) (entities.CartDetail, error) {
	return u.cartSto.GetDetailByID(ctx, cartID)
}

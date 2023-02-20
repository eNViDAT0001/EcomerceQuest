package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/entities"
	"gorm.io/gorm"
)

func (u *cartUseCase) ListCartByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (carts []entities.CartDetail, total int64, err error) {
	total, err = u.cartSto.CountListCartByUserID(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	carts, err = u.cartSto.ListCartByUserID(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}

	return carts, total, err
}

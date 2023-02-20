package cart

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/entities"
)

type UseCase interface {
	GetDetailByID(ctx context.Context, cartID uint) (entities.CartDetail, error)
	ListCartByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (carts []entities.CartDetail, total int64, err error)
	DeleteCart(ctx context.Context, userID uint, cartID uint) error
}

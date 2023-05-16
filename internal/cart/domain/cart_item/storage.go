package cart_item

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item/storage/io"
)

type Storage interface {
	UpsertCartItem(ctx context.Context, providerID uint, input io.CartItemCreateForm) error
	UpdateCartItem(ctx context.Context, itemID uint, cartID uint, quantity int) error
	DeleteCartItem(ctx context.Context, cartID uint, itemID uint) error
}

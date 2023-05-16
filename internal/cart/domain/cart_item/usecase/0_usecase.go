package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item"
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item/storage/io"
)

type cartItemUseCase struct {
	cartItemSto cart_item.Storage
}

func (u *cartItemUseCase) UpsertCartItem(ctx context.Context, providerID uint, input io.CartItemCreateForm) error {
	return u.cartItemSto.UpsertCartItem(ctx, providerID, input)
}

func (u *cartItemUseCase) UpdateCartItem(ctx context.Context, itemID uint, cartID uint, quantity int) error {
	return u.cartItemSto.UpdateCartItem(ctx, itemID, cartID, quantity)
}

func (u *cartItemUseCase) DeleteCartItem(ctx context.Context, cartID uint, itemID uint) error {
	return u.cartItemSto.DeleteCartItem(ctx, cartID, itemID)
}

func NewCartItemUseCase(cartItemSto cart_item.Storage) cart_item.UseCase {
	return &cartItemUseCase{cartItemSto: cartItemSto}
}

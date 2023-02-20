package cart_items

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/domain/cart_item"
)

type cartItemHandler struct {
	cartItemUC cart_item.UseCase
}

func NewCartItemHandler(cartItemUC cart_item.UseCase) cart_item.HttpHandler {
	return &cartItemHandler{cartItemUC: cartItemUC}
}

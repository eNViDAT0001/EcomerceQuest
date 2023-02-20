package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/domain/cart_item"
)

type cartItemStorage struct {
}

func NewCartItemStorage() cart_item.Storage {
	return &cartItemStorage{}
}

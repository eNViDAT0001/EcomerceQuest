package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart"
)

type cartStorage struct {
}

func NewCartStorage() cart.Storage {
	return &cartStorage{}
}

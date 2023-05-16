package usecase

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart"
)

type cartUseCase struct {
	cartSto cart.Storage
}

func NewCartUseCase(cartSto cart.Storage) cart.UseCase {
	return &cartUseCase{cartSto: cartSto}
}

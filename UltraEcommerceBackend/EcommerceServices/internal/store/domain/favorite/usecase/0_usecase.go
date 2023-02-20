package usecase

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/favorite"
)

type favoriteUseCase struct {
	favoriteSto favorite.Storage
}

func NewFavoriteUseCase(
	favoriteSto favorite.Storage,

) favorite.UseCase {
	return &favoriteUseCase{
		favoriteSto: favoriteSto,
	}
}

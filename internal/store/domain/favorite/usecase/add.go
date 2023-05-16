package usecase

import (
	"context"
)

func (u *favoriteUseCase) AddFavorite(ctx context.Context, userID uint, providerID uint) error {
	return u.favoriteSto.AddFavorite(ctx, userID, providerID)
}

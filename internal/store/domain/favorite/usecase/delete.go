package usecase

import (
	"context"
)

func (u *favoriteUseCase) DeleteFavorite(ctx context.Context, userID uint, providerID uint) error {
	return u.DeleteFavorite(ctx, userID, providerID)
}

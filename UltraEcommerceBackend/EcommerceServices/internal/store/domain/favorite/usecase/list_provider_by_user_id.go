package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
	"gorm.io/gorm"
)

func (u *favoriteUseCase) ListFavoriteByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (providers []entities.Provider, total int64, err error) {
	total, err = u.favoriteSto.CountListFavoriteByUserID(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	providers, err = u.favoriteSto.ListFavoriteByUserID(ctx, userID, filter)
	if err != nil {
		return nil, 0, err
	}

	return providers, total, err
}

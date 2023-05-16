package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (f favoriteStorage) CountListFavoriteByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (int64, error) {
	var total int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Provider{}).
		Where("Favorite.user_id = ?", userID).
		Joins("JOIN Favorite ON Favorite.provider_id = Provider.id").
		Joins("JOIN User ON User.id = Favorite.user_id")

	paging_query.SetCountListPagingQuery(&filter, entities.Provider{}.TableName(), query)

	err := query.Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

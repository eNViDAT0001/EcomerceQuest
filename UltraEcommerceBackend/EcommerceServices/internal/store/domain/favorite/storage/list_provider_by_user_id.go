package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

func (f favoriteStorage) ListFavoriteByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) ([]entities.Provider, error) {
	result := make([]entities.Provider, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Provider{}).
		Where("Favorite.user_id = ?", userID).
		Joins("JOIN Favorite ON Favorite.provider_id = Provider.id").
		Joins("JOIN User ON User.id = Favorite.user_id")

	paging_query.SetPagingQuery(&filter, entities.Provider{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

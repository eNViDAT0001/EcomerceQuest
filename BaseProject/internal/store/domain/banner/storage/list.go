package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (b bannerStorage) ListBanner(ctx context.Context, filter paging.ParamsInput) ([]entities.Banner, error) {
	result := make([]entities.Banner, 0)

	db := wrap_gorm.GetDB()

	query := db.Model(entities.Banner{})

	paging_query.SetPagingQuery(&filter, entities.Banner{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

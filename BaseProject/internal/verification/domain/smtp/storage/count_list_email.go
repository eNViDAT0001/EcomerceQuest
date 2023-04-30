package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

func (s smtpStorage) CountListEmail(ctx context.Context, filter paging.ParamsInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Email{}.TableName())

	paging_query.SetCountListPagingQuery(&filter, entities.Email{}.TableName(), query)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

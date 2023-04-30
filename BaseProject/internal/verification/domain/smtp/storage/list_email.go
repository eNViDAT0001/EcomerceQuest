package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
)

func (s smtpStorage) ListEmail(ctx context.Context, filter paging.ParamsInput) ([]entities.Email, error) {
	result := make([]entities.Email, 0)
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Email{}.TableName())
	paging_query.SetPagingQuery(&filter, entities.Email{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

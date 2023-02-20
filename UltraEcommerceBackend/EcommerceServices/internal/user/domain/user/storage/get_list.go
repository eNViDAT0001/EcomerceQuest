package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userStorage) GetUserList(ctx context.Context, input *paging.ParamsInput) ([]*entities.User, error) {
	result := make([]*entities.User, 0)

	db := wrap_gorm.GetDB()

	query := db.Model(entities.User{})

	paging_query.SetPagingQuery(input, entities.User{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

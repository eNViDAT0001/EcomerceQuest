package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"gorm.io/gorm"
)

var wViper = wrap_viper.GetViper()

type requestContactStorage struct {
}

func (r requestContactStorage) UpdateRequestContact(ctx context.Context, id uint, input io.UpdateRequestContactForm) (err error) {
	return wrap_gorm.GetDB().Table(entities.RequestContact{}.TableName()).Where("id = ?", id).Updates(&input).Error
}

func (r requestContactStorage) CreateRequest(ctx context.Context, input io.CreateRequestContact) (id uint, err error) {
	err = wrap_gorm.GetDB().Table(entities.RequestContact{}.TableName()).Create(&input).Error
	return input.ID, err
}

func (r requestContactStorage) ListRequest(ctx context.Context, filter paging.ParamsInput) (requests []entities.RequestContact, total int64, err error) {
	db := wrap_gorm.GetDB()

	query := db.Model(entities.RequestContact{})
	paging_query.SetCountListPagingQuery(&filter, entities.RequestContact{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	result := make([]entities.RequestContact, 0)
	query = db.Model(entities.RequestContact{})
	paging_query.SetPagingQuery(&filter, entities.RequestContact{}.TableName(), query)
	err = query.Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (r requestContactStorage) DeleteRequest(ctx context.Context, id uint) (err error) {
	return wrap_gorm.GetDB().
		Table(entities.RequestContact{}.TableName()).
		Where("id = ?", id).
		Delete(&entities.RequestContact{}).
		Error
}

func NewRequestContactStorage() request_contact.Storage {
	return &requestContactStorage{}
}

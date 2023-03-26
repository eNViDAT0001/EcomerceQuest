package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm/clause"
)

func (s productStorage) CreateSpecification(ctx context.Context, input io.ProductSpecificationCreateForm) (specID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Model(entities.ProductSpecification{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).
		Create(&input).Error
	if err != nil {
		return 0, err
	}
	return input.ID, err
}

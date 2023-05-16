package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

func (s providerStorage) CreateProvider(ctx context.Context, input io.ProviderForm) (io.ProviderForm, error) {
	db := wrap_gorm.GetDB()

	err := db.Table(entities.Provider{}.TableName()).Create(&input).Error

	if err != nil {
		return input, err
	}

	return input, nil
}

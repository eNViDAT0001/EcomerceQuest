package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userStorage) GetUserWithIdentify(ctx context.Context, input io.UserIdentify) (*io.UserPreview, error) {

	var result io.UserPreview
	db := wrap_gorm.GetDB()

	query := db.Model(&entities.User{}).Where("username = ?", input.Username)

	if input.Phone != "" {
		query = query.Or("phone = ?", input.Phone)
	}

	if input.Email != "" {
		query = query.Or("email = ?", input.Email)
	}

	err := query.
		Where("deleted_at IS NULL").
		First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

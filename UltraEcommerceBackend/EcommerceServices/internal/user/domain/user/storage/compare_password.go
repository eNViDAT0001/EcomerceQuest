package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"

	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
	"gorm.io/gorm"
)

func (u userStorage) ComparePassword(ctx context.Context, userID uint32, password string) (io.UserPassword, error) {
	var userPass io.UserPassword
	db := wrap_gorm.GetDB()

	err := db.Model(entities.User{}).
		Select("password, salt").
		Where("id = ?", userID).
		First(&userPass).
		Error

	if err != nil {
		return userPass, err
	}

	ok := entities.User{}.ComparePassword(userPass.Password, password+userPass.Salt)
	if !ok {
		return userPass, gorm.ErrRecordNotFound
	}

	return userPass, nil
}

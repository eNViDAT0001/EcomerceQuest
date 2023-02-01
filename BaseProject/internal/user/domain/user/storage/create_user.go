package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	"time"
)

func (u userStorage) CreateUser(ctx context.Context, input *io.CreateUserInput) (userID uint, err error) {
	db := wrap_gorm.GetDB()
	salt := time.Now().String()
	hashedPassword, err := entities.User{}.HashPassword(input.Password + salt)
	if err != nil {
		return 0, err
	}

	input.Password = hashedPassword
	input.Salt = salt

	newUser := &entities.User{
		Username: &input.Username,
		Password: &input.Password,
		Salt:     &input.Salt,
		Name:     &input.Name,
		Birthday: &input.Birthday,
		Gender:   &input.Gender,
		Email:    &input.Email,
		Phone:    &input.Phone,
		Type:     &input.Type,
	}
	err = db.Create(&newUser).Error

	if err != nil {
		return 0, err
	}
	userID = newUser.ID
	return userID, nil
}

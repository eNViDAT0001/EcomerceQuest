package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
	"gorm.io/gorm"
)

func (u userUseCase) CreateUser(ctx context.Context, input *io.CreateUserInput) (userID uint, err error) {

	user, err := u.userSto.GetUserWithIdentify(ctx, io.UserIdentify{
		Username: input.Username,
		Email:    input.Email,
		Phone:    input.Phone,
	})

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if err != gorm.ErrRecordNotFound {
		if user.Username == input.Username {
			return 0, request.NewConflictError("Username", input.Username, "value exist")
		}
		if user.Email == input.Email {
			return 0, request.NewConflictError("Email", input.Email, "value exist")
		}
		if user.Phone == input.Phone {
			return 0, request.NewConflictError("Phone", input.Phone, "value exist")
		}
	}

	return u.userSto.CreateUser(ctx, input)
}

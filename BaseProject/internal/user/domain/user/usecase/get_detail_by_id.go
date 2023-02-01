package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userUseCase) GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error) {
	return u.userSto.GetUserDetailByID(ctx, ID)
}

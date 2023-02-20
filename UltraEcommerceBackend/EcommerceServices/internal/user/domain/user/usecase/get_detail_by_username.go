package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

func (u userUseCase) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	return u.userSto.GetUserByUsername(ctx, username)
}

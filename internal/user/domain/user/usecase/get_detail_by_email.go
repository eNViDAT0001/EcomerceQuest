package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u userUseCase) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return u.userSto.GetUserByEmail(ctx, email)
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/domain/user/storage/io"
)

func (u userUseCase) ComparePassword(ctx context.Context, userID uint32, password string) (io.UserPassword, error) {
	return u.userSto.ComparePassword(ctx, userID, password)
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
)

func (u userUseCase) ComparePassword(ctx context.Context, userID uint, password string) (io.UserPassword, error) {
	return u.userSto.ComparePassword(ctx, userID, password)
}

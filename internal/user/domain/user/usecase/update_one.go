package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage/io"
)

func (u userUseCase) UpdateUser(ctx context.Context, userID uint32, input *io.UpdateUserInput) error {
	return u.userSto.UpdateUser(ctx, userID, input)
}

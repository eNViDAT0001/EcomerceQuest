package usecase

import "context"

func (u userUseCase) DeleteUserByIDs(ctx context.Context, IDs []uint) error {
	return u.userSto.DeleteUserByIDs(ctx, IDs)
}

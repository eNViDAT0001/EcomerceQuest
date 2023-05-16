package usecase

import "context"

func (u userUseCase) ResetPassword(ctx context.Context, userID uint, newPassword string) error {
	user, err := u.userSto.GetUserDetailByID(ctx, userID)
	if err != nil {
		return err
	}

	return u.userSto.UpdatePassword(ctx, userID, newPassword+*user.Salt)
}

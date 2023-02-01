package usecase

import "context"

func (u userUseCase) SetPassword(ctx context.Context, userID uint32, password string, newPassword string) error {
	userPass, err := u.userSto.ComparePassword(ctx, userID, password)
	if err != nil {
		return err
	}

	return u.userSto.UpdatePassword(ctx, userID, newPassword+userPass.Salt)
}

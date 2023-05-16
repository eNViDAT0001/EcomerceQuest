package usecase

import (
	"context"
)

func (a addressUseCase) DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error {
	return a.addressSto.DeleteAddressByIDs(ctx, userID, addressIDs)
}

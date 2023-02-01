package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
)

func (a addressUseCase) GetAddressesWithUserID(ctx context.Context, userID uint) ([]ioSto.AddressPreview, error) {
	return a.addressSto.GetAddressesWithUserID(ctx, userID)
}

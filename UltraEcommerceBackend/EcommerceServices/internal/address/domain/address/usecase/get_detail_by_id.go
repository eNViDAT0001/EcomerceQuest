package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/storage/io"
)

func (a addressUseCase) GetAddressDetailByID(ctx context.Context, addressID uint) (ioSto.AddressDetail, error) {
	return a.addressSto.GetAddressDetailByID(ctx, addressID)
}

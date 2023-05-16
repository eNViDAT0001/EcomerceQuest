package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a addressUseCase) GetAddressDetailByID(ctx context.Context, addressID uint) (entities.Address, error) {
	return a.addressSto.GetAddressDetailByID(ctx, addressID)
}

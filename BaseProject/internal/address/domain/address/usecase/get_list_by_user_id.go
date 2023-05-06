package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

func (a addressUseCase) GetAddressesWithUserID(ctx context.Context, userID uint) ([]entities.Address, error) {
	return a.addressSto.GetAddressesWithUserID(ctx, userID)
}

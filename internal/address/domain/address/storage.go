package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

type Storage interface {
	GetAddressesWithUserID(ctx context.Context, userID uint) ([]entities.Address, error)
	GetAddressDetailByID(ctx context.Context, addressID uint) (entities.Address, error)
	UpdateAddress(ctx context.Context, addressID uint, input *io.AddressForm) error
	DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error
	CreateAddress(ctx context.Context, input *io.FullAddressForm) error
}

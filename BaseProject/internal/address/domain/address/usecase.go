package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

type UseCase interface {
	GetAddressesWithUserID(ctx context.Context, userID uint) ([]entities.Address, error)
	GetAddressDetailByID(ctx context.Context, addressID uint) (entities.Address, error)
	UpdateAddress(ctx context.Context, addressID uint, input io.UpdateAddressReq) error
	DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error
	CreateAddress(ctx context.Context, input *io.CreateAddressReq) error
}

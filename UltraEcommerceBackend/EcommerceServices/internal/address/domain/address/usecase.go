package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/address/io"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/usecase/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
)

type UseCase interface {
	GetAddressesWithUserID(ctx context.Context, userID uint) ([]ioSto.AddressPreview, error)
	GetAddressDetailByID(ctx context.Context, addressID uint) (ioSto.AddressDetail, error)
	UpdateAddress(ctx context.Context, addressID uint, input io.UpdateAddressReq) error
	DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error
	GetProvinces(ctx context.Context) ([]entities.Province, error)
	GetProvinceByCode(ctx context.Context, code string) (entities.Province, error)
	GetDistrictsWithProvinceCode(ctx context.Context, provinceCode string) ([]ioUC.DistrictRes, error)
	GetWardWithDistrictCode(ctx context.Context, districtCode string) ([]ioUC.WardRes, error)
	CreateAddress(ctx context.Context, input *io.CreateAddressReq) error
}

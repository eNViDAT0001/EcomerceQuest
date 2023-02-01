package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
)

type Storage interface {
	GetAddressesWithUserID(ctx context.Context, userID uint) ([]io.AddressPreview, error)
	GetAddressDetailByID(ctx context.Context, addressID uint) (io.AddressDetail, error)
	UpdateAddress(ctx context.Context, addressID uint, input *io.AddressForm) error
	DeleteAddressByIDs(ctx context.Context, userID uint, addressIDs []uint) error
	GetProvinces(ctx context.Context) ([]entities.Province, error)
	GetProvinceByCode(ctx context.Context, code string) (entities.Province, error)
	GetDistrictsWithProvinceCode(ctx context.Context, provinceCode string) ([]entities.District, error)
	GetWardWithDistrictCode(ctx context.Context, districtCode string) ([]entities.Ward, error)
	CreateAddress(ctx context.Context, input *io.FullAddressForm) error
}

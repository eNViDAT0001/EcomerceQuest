package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/convert"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/io"
)

func (a addressUseCase) GetDistrictsWithProvinceCode(ctx context.Context, provinceCode string) ([]io.DistrictRes, error) {
	districts, err := a.addressSto.GetDistrictsWithProvinceCode(ctx, provinceCode)
	if err != nil {
		return nil, err
	}

	result, err := convert.DistrictArrayEntityToDistrictArrayRes(districts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

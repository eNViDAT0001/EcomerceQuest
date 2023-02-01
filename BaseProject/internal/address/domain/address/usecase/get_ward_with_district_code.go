package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/convert"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/io"
)

func (a addressUseCase) GetWardWithDistrictCode(ctx context.Context, districtCode string) ([]io.WardRes, error) {
	wards, err := a.addressSto.GetWardWithDistrictCode(ctx, districtCode)
	if err != nil {
		return nil, err
	}

	result, err := convert.WardArrayEntityToWardArrayRes(wards)
	if err != nil {
		return nil, err
	}

	return result, nil
}

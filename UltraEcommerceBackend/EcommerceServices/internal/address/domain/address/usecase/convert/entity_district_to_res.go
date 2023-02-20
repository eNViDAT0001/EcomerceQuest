package convert

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/usecase/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
	"github.com/jinzhu/copier"
)

func DistrictEntityToDistrictRes(DistrictEntity *entities.District) (io.DistrictRes, error) {
	var result io.DistrictRes
	err := copier.Copy(&result, DistrictEntity)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DistrictArrayEntityToDistrictArrayRes(DistrictEntity []entities.District) ([]io.DistrictRes, error) {
	result := make([]io.DistrictRes, 0)

	err := copier.Copy(&result, DistrictEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

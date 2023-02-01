package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
	"github.com/jinzhu/copier"
)

func ProvinceEntityToProvinceRes(ProvinceEntity *entities.Province) (*io.ProvinceRes, error) {
	var result io.ProvinceRes
	err := copier.Copy(&result, ProvinceEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ProvinceArrayEntityToProvinceArrayRes(ProvinceEntity []entities.Province) ([]*io.ProvinceRes, error) {
	result := make([]*io.ProvinceRes, 0)

	err := copier.Copy(&result, ProvinceEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/entities"
	"github.com/jinzhu/copier"
)

func WardEntityToWardRes(WardEntity *entities.Ward) (*io.WardRes, error) {
	var result io.WardRes
	err := copier.Copy(&result, WardEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func WardArrayEntityToWardArrayRes(WardEntity []entities.Ward) ([]io.WardRes, error) {
	result := make([]io.WardRes, 0)

	err := copier.Copy(&result, WardEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

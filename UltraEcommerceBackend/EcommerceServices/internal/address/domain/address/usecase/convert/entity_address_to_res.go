package convert

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/domain/address/usecase/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/address/entities"
	"github.com/jinzhu/copier"
)

func AddressEntityToAddressRes(addressEntity *entities.Address) (*io.AddressRes, error) {
	var result io.AddressRes
	err := copier.Copy(&result, addressEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func AddressArrayEntityToAddressArrayRes(addressEntity []entities.Address) ([]*io.AddressRes, error) {
	result := make([]*io.AddressRes, 0)

	err := copier.Copy(&result, addressEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

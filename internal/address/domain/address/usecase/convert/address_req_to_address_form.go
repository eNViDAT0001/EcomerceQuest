package convert

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage/io"
	"github.com/jinzhu/copier"
)

func CreateAddressReqToFullAddressForm(addressEntity *io.CreateAddressReq) (*ioSto.FullAddressForm, error) {
	var result ioSto.FullAddressForm
	err := copier.Copy(&result, addressEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func UpdateAddressReqToAddressForm(addressEntity *io.UpdateAddressReq) (*ioSto.AddressForm, error) {
	var result ioSto.AddressForm
	err := copier.Copy(&result, addressEntity)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

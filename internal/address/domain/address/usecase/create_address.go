package usecase

import (
	"context"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/convert"
)

func (a addressUseCase) CreateAddress(ctx context.Context, input *ioHandler.CreateAddressReq) error {
	inputReq, err := convert.CreateAddressReqToFullAddressForm(input)
	if err != nil {
		return err
	}
	err = a.addressSto.CreateAddress(ctx, inputReq)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"context"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/address/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase/convert"
)

func (a addressUseCase) UpdateAddress(ctx context.Context, addressID uint, input ioHandler.UpdateAddressReq) error {
	inputReq, err := convert.UpdateAddressReqToAddressForm(&input)
	if err != nil {
		return err
	}
	return a.addressSto.UpdateAddress(ctx, addressID, inputReq)
}

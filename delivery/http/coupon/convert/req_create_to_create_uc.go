package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/jinzhu/copier"
)

func CreateReqToCreateCouponInput(input *ioHttpHandler.CouponCreateReq) (ioSto.CouponCreateForm, error) {
	var result ioSto.CouponCreateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateReqToCreateCouponProductsInput(input []ioHttpHandler.CouponDetailCreateReq) ([]ioSto.CouponDetailCreateForm, error) {
	var result []ioSto.CouponDetailCreateForm
	for _, req := range input {
		var temp ioSto.CouponDetailCreateForm
		err := copier.Copy(&temp, &req)
		if err != nil {
			return nil, err
		}

		result = append(result, temp)
	}
	return result, nil
}

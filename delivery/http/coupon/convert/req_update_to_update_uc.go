package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/coupon/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateReqToUpdateCouponInput(input *ioHttpHandler.CouponUpdateReq) (ioSto.CouponUpdateForm, []ioSto.CouponDetailCreateForm, error) {
	var result ioSto.CouponUpdateForm
	var products []ioSto.CouponDetailCreateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, products, err
	}
	for _, req := range input.ProductsIN {
		var product ioSto.CouponDetailCreateForm
		err = copier.Copy(&product, &req)
		if err != nil {
			return result, products, err
		}
		products = append(products, product)
	}

	return result, products, nil
}

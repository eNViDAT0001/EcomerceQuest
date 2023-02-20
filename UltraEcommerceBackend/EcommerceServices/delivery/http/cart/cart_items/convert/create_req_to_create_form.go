package convert

import (
	ioHTTPHandler "github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/cart/cart_items/io"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/cart/domain/cart_item/storage/io"
	"github.com/jinzhu/copier"
)

func CreateReqToCreateCartItemsFormInput(input *ioHTTPHandler.CartItemCreateReq) (ioSto.CartItemCreateForm, error) {
	var result ioSto.CartItemCreateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}

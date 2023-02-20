package convert

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/order/order/io"
	io2 "github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/domain/order/storage/io"
	io3 "github.com/eNViDAT0001/Thesis/Ecommerce/internal/order/domain/order_item/storage/io"
	"github.com/jinzhu/copier"
)

func CreateReqToCreateOrderFormInput(input *io.CreateOrderReq) (io2.CreateOrderForm, []io3.CreateOrderItemForm, error) {
	var order io2.CreateOrderForm
	var items []io3.CreateOrderItemForm
	err := copier.Copy(&order, &input)
	if err != nil {
		return order, nil, err
	}
	err = copier.Copy(&items, &input.Items)
	if err != nil {
		return order, nil, err
	}

	return order, items, nil
}

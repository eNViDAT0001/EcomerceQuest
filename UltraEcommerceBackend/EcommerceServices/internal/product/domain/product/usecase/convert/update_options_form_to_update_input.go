package convert

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"github.com/jinzhu/copier"
)

func OptionsUpdateFormToOptionsUpdateInput(input *io.ProductOptionUpdateForm) (io.ProductOptionUpdateInput, error) {
	var result io.ProductOptionUpdateInput
	err := copier.Copy(&result, input)
	if err != nil {
		return result, err
	}
	return result, nil
}

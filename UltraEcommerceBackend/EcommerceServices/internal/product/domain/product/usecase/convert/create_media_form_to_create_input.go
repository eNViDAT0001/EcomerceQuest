package convert

import (
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/usecase/io"
	"github.com/jinzhu/copier"
)

func MediaCreateFormToMediaCreateInput(input []ioUC.CreateProductMediaForm) ([]ioSto.CreateProductMedia, error) {
	var result []ioSto.CreateProductMedia
	err := copier.Copy(&result, input)
	if err != nil {
		return result, err
	}
	return result, nil
}

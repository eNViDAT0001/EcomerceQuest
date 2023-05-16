package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateOptionsReqToUpdateOptionsForm(input *ioHttpHandler.ProductOptionsUpdateReq) ([]ioSto.ProductOptionUpdateForm, error) {
	var result []ioSto.ProductOptionUpdateForm
	err := copier.Copy(&result, &input.Options)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateOptionsReqToCreateOptionsForm(input *ioHttpHandler.ProductOptionsCreateReq) ([]ioSto.ProductOptionCreateForm, error) {
	var result []ioSto.ProductOptionCreateForm
	err := copier.Copy(&result, &input.Options)
	if err != nil {
		return result, err
	}
	return result, nil
}

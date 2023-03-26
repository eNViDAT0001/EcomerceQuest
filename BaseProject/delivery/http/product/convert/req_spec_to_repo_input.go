package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"github.com/jinzhu/copier"
)

func UpdateSpecificationReqToUpdateSpecificationForm(input *ioHttpHandler.ProductSpecificationUpdateReq) (ioSto.ProductSpecificationUpdateForm, error) {
	var result ioSto.ProductSpecificationUpdateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateSpecificationReqToCreateSpecificationForm(input *ioHttpHandler.SpecificationCreate) (ioUC.SpecificationCreateForm, error) {
	var result ioUC.SpecificationCreateForm
	var spec ioSto.ProductSpecificationCreateForm
	var options []ioSto.ProductOptionCreateForm

	err := copier.Copy(&spec, &input.Specification)
	if err != nil {
		return result, err
	}
	err = copier.Copy(&options, &input.Options)
	if err != nil {
		return result, err
	}

	result.Specification = spec
	result.Options = options

	return result, nil
}

package usecase

import (
	"context"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"gorm.io/gorm"
)

func (u *productUseCase) CreateSpecification(ctx context.Context, input ioUC.SpecificationCreateForm) (specID uint, err error) {
	if len(input.Options) < 1 {
		return 0, gorm.ErrEmptySlice
	}

	specID, err = u.productSto.CreateSpecification(ctx, input.Specification)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(input.Options); i++ {
		input.Options[i].SpecificationID = specID
		input.Options[i].ProductID = input.Specification.ProductID
	}
	err = u.productSto.CreateProductOptions(ctx, input.Options)
	if err != nil {
		return 0, err
	}
	return specID, nil
}

package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/convert"
)

func (u *productUseCase) UpdateProductOptions(ctx context.Context, options []ioSto.ProductOptionUpdateForm) error {
	for _, option := range options {
		inputSto, err := convert.OptionsUpdateFormToOptionsUpdateInput(&option)
		err = u.productSto.UpdateProductOptions(ctx, option.ID, inputSto)
		if err != nil {
			return err
		}
	}
	return nil
}

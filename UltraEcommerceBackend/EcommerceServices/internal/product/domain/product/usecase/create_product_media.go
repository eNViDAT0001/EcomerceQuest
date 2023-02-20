package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/usecase/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/usecase/io"
)

func (u *productUseCase) CreateProductMedia(ctx context.Context, media []io.CreateProductMediaForm) error {
	ioSto, err := convert.MediaCreateFormToMediaCreateInput(media)
	if err != nil {
		return err
	}
	return u.productSto.CreateProductMedia(ctx, ioSto)
}

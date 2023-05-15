package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"gorm.io/gorm"
)

func (u *productUseCase) ListProductPreviewByBannerID(ctx context.Context, input ioSto.ListProductInput) (items []ioSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByBannerID(ctx, input.BannerID)
	if err != nil {
		return nil, 0, err
	}

	input.ProductIDs = productIDs
	total, err = u.productSto.ListCountProductsPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err := u.productSto.ListProductsPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}
func (u *productUseCase) ListProductPreviewNotInBannerID(ctx context.Context, input ioSto.ListProductInput) (items []ioSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByNotInBannerID(ctx, input.BannerID)
	if err != nil {
		return nil, 0, err
	}

	input.ProductIDs = productIDs
	total, err = u.productSto.ListCountProductsPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err := u.productSto.ListProductsPreview(ctx, input)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

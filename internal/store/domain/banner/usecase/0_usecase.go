package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product"
	ioProductSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	productEntities "github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

type bannerUseCase struct {
	bannerSto  banner.Storage
	productSto product.Storage
}

func (u bannerUseCase) CreateBanner(ctx context.Context, input io.BannerCreateForm, productIDs []uint) (BannerID uint, err error) {
	return u.bannerSto.CreateBanner(ctx, input, productIDs)
}

func (u bannerUseCase) GetBannerByID(ctx context.Context, bannerID uint) (io.BannerDetail, error) {
	return u.bannerSto.GetBannerByID(ctx, bannerID)
}
func (u bannerUseCase) GetBannerDetailByID(ctx context.Context, bannerID uint) (entities.Banner, error) {
	return u.bannerSto.GetBannerDetailByID(ctx, bannerID)
}

func (u bannerUseCase) UpdateBanner(ctx context.Context, bannerID uint, input io.BannerUpdateForm, productIDsIN []uint, productIDsOUT []uint) error {
	return u.bannerSto.UpdateBanner(ctx, bannerID, input, productIDsIN, productIDsOUT)
}

func (u bannerUseCase) DeleteBannerByIDs(ctx context.Context, bannerID []uint) error {
	return u.bannerSto.DeleteBannerByIDs(ctx, bannerID)
}

func (u bannerUseCase) ListBanner(ctx context.Context, filter paging.ParamsInput) (banners []entities.Banner, total int64, err error) {
	total, err = u.bannerSto.CountListBanner(ctx, filter, 0)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	banners, err = u.bannerSto.ListBanner(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return banners, total, err
}

func (u bannerUseCase) ListProductPreviewByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioProductSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByBannerID(ctx, bannerID)
	if err != nil {
		return nil, 0, err
	}
	if len(productIDs) < 1 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	productFilter := ioProductSto.ListProductInput{
		ProductIDs: productIDs,
		Paging:     filter,
	}
	total, err = u.productSto.ListCountProductsPreview(ctx, productFilter)

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err = u.productSto.ListProductsPreview(ctx, productFilter)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

func (u bannerUseCase) ListProductByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []productEntities.Product, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByBannerID(ctx, bannerID)
	if err != nil {
		return nil, 0, err
	}
	if len(productIDs) < 1 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	productFilter := ioProductSto.ListProductInput{
		ProductIDs: productIDs,
		Paging:     filter,
	}
	total, err = u.productSto.ListCountProduct(ctx, productFilter)

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err = u.productSto.ListProduct(ctx, productFilter)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}
func (u bannerUseCase) ListProductPreviewNotInBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioProductSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByNotInBannerID(ctx, bannerID)
	if err != nil {
		return nil, 0, err
	}

	productFilter := ioProductSto.ListProductInput{
		ProductIDs: productIDs,
		Paging:     filter,
	}
	total, err = u.productSto.ListCountProductsPreview(ctx, productFilter)

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err = u.productSto.ListProductsPreview(ctx, productFilter)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}
func (u bannerUseCase) ListProductNotINBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []productEntities.Product, total int64, err error) {
	productIDs, err := u.bannerSto.ProductIDsByNotInBannerID(ctx, bannerID)
	if err != nil {
		return nil, 0, err
	}

	productFilter := ioProductSto.ListProductInput{
		ProductIDs: productIDs,
		Paging:     filter,
	}
	total, err = u.productSto.ListCountProduct(ctx, productFilter)

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	products, err = u.productSto.ListProduct(ctx, productFilter)
	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}
func NewBannerUseCase(
	bannerSto banner.Storage,
	productSto product.Storage,
) banner.UseCase {
	return &bannerUseCase{
		bannerSto:  bannerSto,
		productSto: productSto,
	}
}

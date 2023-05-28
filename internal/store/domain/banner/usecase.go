package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type UseCase interface {
	CreateBanner(ctx context.Context, input io.BannerCreateForm, productIDs []uint) (BannerID uint, err error)
	GetBannerByID(ctx context.Context, bannerID uint) (io.BannerDetail, error)
	GetBannerDetailByID(ctx context.Context, bannerID uint) (entities.Banner, error)
	UpdateBanner(ctx context.Context, bannerID uint, input io.BannerUpdateForm, productIDsIN []uint, productIDsOUT []uint) error
	DeleteBannerByIDs(ctx context.Context, bannerID []uint) error
	ListBanner(ctx context.Context, filter paging.ParamsInput) (banners []entities.Banner, total int64, err error)
	ListProductPreviewByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)

	ListProductPreviewNotInBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductNotINBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)
}

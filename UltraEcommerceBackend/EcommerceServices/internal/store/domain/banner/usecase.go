package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	productEntities "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/entities"
)

type UseCase interface {
	CreateBanner(ctx context.Context, input io.BannerCreateForm, productIDs []uint) (BannerID uint, err error)
	GetBannerByID(ctx context.Context, bannerID uint) (io.BannerDetail, error)
	UpdateBanner(ctx context.Context, bannerID uint, input io.BannerUpdateForm, productIDsIN []uint, productIDsOUT []uint) error
	DeleteBannerByIDs(ctx context.Context, bannerID []uint) error
	ListBanner(ctx context.Context, filter paging.ParamsInput) (banners []entities.Banner, total int64, err error)
	ListProductPreviewByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) (products []productEntities.Product, total int64, err error)
}

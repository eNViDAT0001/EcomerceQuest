package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type Storage interface {
	CreateBanner(ctx context.Context, input io.BannerCreateForm, productIDs []uint) (BannerID uint, err error)
	GetBannerByID(ctx context.Context, bannerID uint) (io.BannerDetail, error)
	UpdateBanner(ctx context.Context, bannerID uint, input io.BannerUpdateForm, productIDsIN []uint, productIDsOUT []uint) error
	DeleteBannerByIDs(ctx context.Context, bannerID []uint) error
	ListBanner(ctx context.Context, filter paging.ParamsInput) ([]entities.Banner, error)
	CountListBanner(ctx context.Context, filter paging.ParamsInput, bannerID uint) (total int64, err error)
	ListProductIDsByBannerID(ctx context.Context, bannerID uint, filter paging.ParamsInput) ([]uint, error)
	ProductIDsByNotInBannerID(ctx context.Context, bannerID uint) ([]uint, error)
	ProductIDsByBannerID(ctx context.Context, bannerID uint) ([]uint, error)
}

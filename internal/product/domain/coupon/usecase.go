package Coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

type UseCase interface {
	CreateCoupon(ctx context.Context, input io.CouponCreateForm, productIDs []uint) (CouponID uint, err error)
	GetCouponByID(ctx context.Context, CouponID uint) (io.CouponDetail, error)
	UpdateCoupon(ctx context.Context, CouponID uint, input io.CouponUpdateForm, productIDsIN []uint, productIDsOUT []uint) error
	DeleteCouponByIDs(ctx context.Context, CouponID []uint) error

	ListCoupon(ctx context.Context, filter paging.ParamsInput) (Coupons []entities.Coupon, total int64, err error)

	ListProductPreviewByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)

	ListProductPreviewNotInCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductNotINCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)

	ValidateCouponByProductIDs(ctx context.Context, CouponCode uint, productIDs []uint) (bool, error)
}

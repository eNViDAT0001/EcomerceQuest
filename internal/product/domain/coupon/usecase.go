package Coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/usecase/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

type UseCase interface {
	CreateCoupon(ctx context.Context, input io.CouponCreateForm, products []io.CouponDetailCreateForm) (couponID uint, err error)
	GetCouponByID(ctx context.Context, couponID uint) (io.CouponDetail, error)
	UpdateCoupon(ctx context.Context, couponID uint, input io.CouponUpdateForm, productsIN []io.CouponDetailCreateForm, productIDsOUT []uint) error
	DeleteCouponByIDs(ctx context.Context, userID uint, couponID []uint) error

	ListCoupon(ctx context.Context, filter paging.ParamsInput) (Coupons []entities.Coupon, total int64, err error)

	ListProductPreviewByCouponID(ctx context.Context, couponID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)

	ListProductPreviewNotInCouponID(ctx context.Context, couponID uint, filter paging.ParamsInput) (products []ioSto.ProductPreviewItem, total int64, err error)
	ListProductNotINCouponID(ctx context.Context, couponID uint, filter paging.ParamsInput) (products []ioSto.ProductWithQuantities, total int64, err error)

	ValidateCouponByProductIDs(ctx context.Context, CouponCode string, productIDs []uint) ([]ioUC.ValidatedProduct, error)
}

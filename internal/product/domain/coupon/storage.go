package Coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/gorm"
)

type Storage interface {
	CreateCoupon(ctx context.Context, input io.CouponCreateForm, products []io.CouponDetailCreateForm) (CouponID uint, err error)
	GetCouponByID(ctx context.Context, couponID uint) (io.CouponDetail, error)
	UpdateCoupon(ctx context.Context, couponID uint, input io.CouponUpdateForm, productsIN []io.CouponDetailCreateForm, productIDsOUT []uint) error
	DeleteCouponByIDs(ctx context.Context, userID uint, couponIDs []uint) error
	ListCoupon(ctx context.Context, filter paging.ParamsInput) ([]entities.Coupon, error)
	CountListCoupon(ctx context.Context, filter paging.ParamsInput, couponID uint) (total int64, err error)
	ListProductIDsByCouponID(ctx context.Context, couponID uint, filter paging.ParamsInput) ([]uint, error)
	ProductIDsByNotInCouponID(ctx context.Context, couponID uint) ([]uint, error)
	ProductIDsByCouponID(ctx context.Context, couponID uint) ([]uint, error)

	ValidateCouponByProductIDs(ctx context.Context, couponCode string, productID uint) (entities.Coupon, error)
	UseCouponByProductIDsWithGorm(ctx context.Context, db *gorm.DB, couponCode []string, productIDs []uint) error
}

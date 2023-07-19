package Coupon

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type Storage interface {
	CreateCoupon(ctx context.Context, input io.CouponCreateForm, productIDs []uint) (CouponID uint, err error)
	GetCouponByID(ctx context.Context, CouponID uint) (io.CouponDetail, error)
	GetCouponDetailByID(ctx context.Context, CouponID uint) (entities.Coupon, error)
	UpdateCoupon(ctx context.Context, CouponID uint, input io.CouponUpdateForm, productIDsIN []uint, productIDsOUT []uint) error
	DeleteCouponByIDs(ctx context.Context, CouponID []uint) error
	ListCoupon(ctx context.Context, filter paging.ParamsInput) ([]entities.Coupon, error)
	CountListCoupon(ctx context.Context, filter paging.ParamsInput, CouponID uint) (total int64, err error)
	ListProductIDsByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) ([]uint, error)
	ProductIDsByNotInCouponID(ctx context.Context, CouponID uint) ([]uint, error)
	ProductIDsByCouponID(ctx context.Context, CouponID uint) ([]uint, error)
}

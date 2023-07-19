package usecase

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product"
	ioProductSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/product/coupon"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/product/coupon/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

type couponUseCase struct {
	couponSto  coupon.Storage
	productSto product.Storage
}

func (u couponUseCase) CreateCoupon(ctx context.Context, input io.CouponCreateForm, productIDs []uint) (CouponID uint, err error) {
	return u.couponSto.CreateCoupon(ctx, input, productIDs)
}

func (u couponUseCase) GetCouponByID(ctx context.Context, CouponID uint) (io.CouponDetail, error) {
	return u.couponSto.GetCouponByID(ctx, CouponID)
}
func (u couponUseCase) GetCouponDetailByID(ctx context.Context, CouponID uint) (entities.Coupon, error) {
	return u.couponSto.GetCouponDetailByID(ctx, CouponID)
}

func (u couponUseCase) UpdateCoupon(ctx context.Context, CouponID uint, input io.CouponUpdateForm, productIDsIN []uint, productIDsOUT []uint) error {
	return u.couponSto.UpdateCoupon(ctx, CouponID, input, productIDsIN, productIDsOUT)
}

func (u couponUseCase) DeleteCouponByIDs(ctx context.Context, CouponID []uint) error {
	return u.couponSto.DeleteCouponByIDs(ctx, CouponID)
}

func (u couponUseCase) ListCoupon(ctx context.Context, filter paging.ParamsInput) (Coupons []entities.Coupon, total int64, err error) {
	total, err = u.couponSto.CountListCoupon(ctx, filter, 0)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	Coupons, err = u.couponSto.ListCoupon(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return Coupons, total, err
}

func (u couponUseCase) ListProductPreviewByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioProductSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.couponSto.ProductIDsByCouponID(ctx, CouponID)
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

func (u couponUseCase) ListProductByCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioProductSto.ProductWithQuantities, total int64, err error) {
	productIDs, err := u.couponSto.ProductIDsByCouponID(ctx, CouponID)
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
func (u couponUseCase) ListProductPreviewNotInCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioProductSto.ProductPreviewItem, total int64, err error) {
	productIDs, err := u.couponSto.ProductIDsByNotInCouponID(ctx, CouponID)
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
func (u couponUseCase) ListProductNotINCouponID(ctx context.Context, CouponID uint, filter paging.ParamsInput) (products []ioProductSto.ProductWithQuantities, total int64, err error) {
	productIDs, err := u.couponSto.ProductIDsByNotInCouponID(ctx, CouponID)
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
func NewCouponUseCase(
	couponSto coupon.Storage,
	productSto product.Storage,
) coupon.UseCase {
	return &couponUseCase{
		couponSto:  couponSto,
		productSto: productSto,
	}
}

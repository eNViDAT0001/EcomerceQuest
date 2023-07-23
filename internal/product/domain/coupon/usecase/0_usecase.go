package usecase

import (
	"context"
	coupon "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/coupon/usecase/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"

	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product"
	ioProductSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"

	"gorm.io/gorm"
)

type couponUseCase struct {
	couponSto  coupon.Storage
	productSto product.Storage
}

func (u couponUseCase) ValidateCouponByProductIDs(ctx context.Context, CouponCode string, productIDs []uint) ([]ioUC.ValidatedProduct, error) {
	var result []ioUC.ValidatedProduct
	for _, id := range productIDs {
		cp, err := u.couponSto.ValidateCouponByProductIDs(ctx, CouponCode, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if err != gorm.ErrRecordNotFound {
			result = append(result, ioUC.ValidatedProduct{
				ProductID: id,
				Total:     cp.Fixed,
			})
		}
	}
	return result, nil
}

func (u couponUseCase) CreateCoupon(ctx context.Context, input io.CouponCreateForm, products []io.CouponDetailCreateForm) (couponID uint, err error) {
	return u.couponSto.CreateCoupon(ctx, input, products)
}

func (u couponUseCase) GetCouponByID(ctx context.Context, couponID uint) (io.CouponDetail, error) {
	return u.couponSto.GetCouponByID(ctx, couponID)
}

func (u couponUseCase) UpdateCoupon(ctx context.Context, couponID uint, input io.CouponUpdateForm, productsIN []io.CouponDetailCreateForm, productIDsOUT []uint) error {
	return u.couponSto.UpdateCoupon(ctx, couponID, input, productsIN, productIDsOUT)
}

func (u couponUseCase) DeleteCouponByIDs(ctx context.Context, userID uint, couponID []uint) error {
	return u.couponSto.DeleteCouponByIDs(ctx, userID, couponID)
}

func (u couponUseCase) ListCoupon(ctx context.Context, filter paging.ParamsInput) (coupons []entities.Coupon, total int64, err error) {
	total, err = u.couponSto.CountListCoupon(ctx, filter, 0)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	coupons, err = u.couponSto.ListCoupon(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return coupons, total, err
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

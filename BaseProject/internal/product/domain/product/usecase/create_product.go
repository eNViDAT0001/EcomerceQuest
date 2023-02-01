package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/convert"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
)

func (u *productUseCase) CreateProduct(ctx context.Context, productDetail ioUC.ProductDetailCreateForm) (productID uint, err error) {
	productID, err = u.productSto.CreateProduct(ctx, productDetail.Product)
	if err != nil {
		return 0, err
	}

	if len(productDetail.Media) > 0 {
		for i, _ := range productDetail.Media {
			productDetail.Media[i].ProductID = productID
		}
		mediaSto, err := convert.MediaCreateFormToMediaCreateInput(productDetail.Media)
		if err != nil {
			return productID, err
		}

		err = u.productSto.CreateProductMedia(ctx, mediaSto)
		if err != nil {
			return productID, err
		}
	}
	if len(productDetail.Specifications) > 0 {
		for i, _ := range productDetail.Specifications {
			productDetail.Specifications[i].Specification.ProductID = productID
		}

		err = u.CreateSpecificationTree(ctx, productDetail.Specifications)
	}
	if len(productDetail.Descriptions) > 0 {
		for i, _ := range productDetail.Descriptions {
			productDetail.Descriptions[i].ProductID = productID
		}

		err = u.productSto.CreateBulkProductDescriptions(ctx, productDetail.Descriptions)
	}

	return productID, nil
}

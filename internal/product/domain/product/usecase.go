package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"mime/multipart"
)

type UseCase interface {
	ListProductsPreview(ctx context.Context, input io.ListProductInput) (items []io.ProductPreviewItem, total int64, err error)
	ListProduct(ctx context.Context, input io.ListProductInput) (items []io.ProductWithQuantities, total int64, err error)
	ListProductPreviewByBannerID(ctx context.Context, input io.ListProductInput) (items []io.ProductPreviewItem, total int64, err error)
	ListProductPreviewNotInBannerID(ctx context.Context, input io.ListProductInput) (items []io.ProductPreviewItem, total int64, err error)
	ListRecommendedProductsPreview(ctx context.Context, input io.ListRecommendedProductInput) (items []entities.ProductPreview, total int64, err error)
	ListProductQuantity(ctx context.Context, input io.ListProductInput) (items []io.ProductQuantity, total int64, err error)
	ListProductMediaByProductID(ctx context.Context, productID uint) ([]entities.ProductMedia, error)

	GetSpecificationTreeByProductID(ctx context.Context, productID uint) ([]io.ProductSpecificationFullDetail, error)
	GetProductDescriptionsByProductID(ctx context.Context, productID uint) ([]entities.ProductDescriptions, error)
	GetProductInfoByID(ctx context.Context, productID uint) (entities.Product, error)

	UpdateFullProduct(ctx context.Context, productID uint, product io.ProductFullUpdateForm) error
	UpdateProduct(ctx context.Context, productID uint, product io.ProductUpdateForm) error
	UpdateProductOptions(ctx context.Context, options []io.ProductOptionUpdateForm) error
	UpdateProductSpecification(ctx context.Context, specID uint, specifications io.ProductSpecificationUpdateForm) error
	UpdateProductDescriptions(ctx context.Context, descriptionsID uint, descriptions ioUC.ProductDescriptionsUpdate) error

	CreateProduct(ctx context.Context, productDetail ioUC.ProductDetailCreateForm) (productID uint, err error)
	CreateProductMedia(ctx context.Context, media []ioUC.CreateProductMediaForm) error
	CreateProductMediaWithFiles(ctx context.Context, productID uint, files []*multipart.FileHeader) error
	CreateDescriptions(ctx context.Context, input ioUC.ProductDescriptionsWithFileCreate) (newID uint, err error)
	CreateProductOptions(ctx context.Context, input []io.ProductOptionCreateForm) error
	CreateSpecification(ctx context.Context, input ioUC.SpecificationCreateForm) (specID uint, err error)

	DeleteProductByIDs(ctx context.Context, IDs []uint) error
	DeleteProductByID(ctx context.Context, ID uint) error
	DeleteElementByIDs(ctx context.Context, ID uint, descriptionsIDs []uint, mediaIDs []uint, optionIDs []uint) error
}

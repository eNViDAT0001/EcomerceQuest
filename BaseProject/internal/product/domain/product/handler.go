package product

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListProductsPreview() func(ctx *gin.Context)
	ListProduct() func(ctx *gin.Context)
	ListProductPreviewWithCategoryID() func(ctx *gin.Context)
	ListProductWithCategoryID() func(ctx *gin.Context)

	GetProductMediaByProductID() func(ctx *gin.Context)

	GetSpecificationTreeByProductID() func(ctx *gin.Context)
	GetProductDescriptionsByProductID() func(ctx *gin.Context)
	GetProductInfoByID() func(ctx *gin.Context)

	UpdateProduct() func(ctx *gin.Context)
	UpdateProductOptions() func(ctx *gin.Context)
	UpdateProductSpecification() func(ctx *gin.Context)
	UpdateDescriptions() func(ctx *gin.Context)

	CreateProduct() func(ctx *gin.Context)
	CreateProductMedia() func(ctx *gin.Context)
	CreateDescriptions() func(ctx *gin.Context)
	CreateProductOptions() func(ctx *gin.Context)
	CreateSpecification() func(ctx *gin.Context)
	CreateSpecificationTree() func(ctx *gin.Context)

	DeleteProductByIDs() func(ctx *gin.Context)
	DeleteProductByID() func(ctx *gin.Context)
}

package product

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/file_storage/domain/media"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category"
)

type productHandler struct {
	productUC  product.UseCase
	mediaUC    media.UseCase
	categoryUC category.UseCase
}

func NewProductHandler(productUC product.UseCase, mediaUC media.UseCase, categoryUC category.UseCase) product.HttpHandler {
	return &productHandler{productUC: productUC, mediaUC: mediaUC, categoryUC: categoryUC}
}

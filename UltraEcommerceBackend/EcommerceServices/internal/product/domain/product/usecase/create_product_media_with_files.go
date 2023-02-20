package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_cloudinary"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/product/storage/io"
	"mime/multipart"
	"strings"
)

func (u *productUseCase) CreateProductMediaWithFiles(ctx context.Context, productID uint, files []*multipart.FileHeader) error {
	uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, wrap_cloudinary.Product)
	if err != nil {
		return err
	}

	mediaStore := make([]ioSto.CreateProductMedia, 0)
	for _, file := range uploadedFiles {
		ioMediaCreate := ioSto.CreateProductMedia{
			ProductID: productID,
			PublicID:  file.PublicID,
			MediaPath: file.URL,
			MediaType: strings.ToUpper(file.ResourceType),
		}
		mediaStore = append(mediaStore, ioMediaCreate)
	}
	return u.productSto.CreateProductMedia(ctx, mediaStore)
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_cloudinary"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"mime/multipart"
)

func (u *productUseCase) CreateDescriptionsWithFiles(ctx context.Context, input []ioUC.ProductDescriptionsWithFileCreate) error {
	for _, v := range input {
		files := []*multipart.FileHeader{v.File}
		uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, wrap_cloudinary.ProductDescriptions)
		if err != nil {
			return err
		}
		descriptions := io.ProductDescriptionsCreateForm{
			ProductID:        v.ProductID,
			Name:             v.Name,
			PublicID:         uploadedFiles[0].PublicID,
			DescriptionsPath: uploadedFiles[0].URL,
		}

		_, err = u.productSto.CreateProductDescriptions(ctx, descriptions)
		if err != nil {
			return err
		}
	}

	return nil
}

package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_cloudinary"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"mime/multipart"
)

func (u *productUseCase) UpdateProductDescriptions(ctx context.Context, descriptionsID uint, descriptions ioUC.ProductDescriptionsUpdate) error {

	var inputSto ioSto.ProductDescriptionsUpdateForm
	if descriptions.File != nil {
		files := []*multipart.FileHeader{descriptions.File}
		uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, wrap_cloudinary.ProductDescriptions)
		if err != nil {
			return err
		}
		inputSto.DescriptionsPath = uploadedFiles[0].URL
		inputSto.PublicID = uploadedFiles[0].PublicID
	}

	if descriptions.Name != nil {
		inputSto.Name = *descriptions.Name
	}

	return u.productSto.UpdateProductDescriptions(ctx, descriptionsID, inputSto)
}

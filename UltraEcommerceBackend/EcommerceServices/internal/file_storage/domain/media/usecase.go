package media

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_cloudinary"
	"mime/multipart"
)

type UseCase interface {
	UploadMedia(ctx context.Context, files []*multipart.FileHeader, folder wrap_cloudinary.MediaFolderType) ([]*uploader.UploadResult, error)
	DeleteMedia(ctx context.Context, publicID string) (*uploader.DestroyResult, error)
}

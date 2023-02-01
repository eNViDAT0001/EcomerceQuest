package storage

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	wrap_cloudinary2 "github.com/eNViDAT0001/Thesis/Backend/external/wrap_cloudinary"
	"mime/multipart"
)

func (m mediaStorage) UploadMedia(ctx context.Context, files []*multipart.FileHeader, folder wrap_cloudinary2.MediaFolderType) ([]*uploader.UploadResult, error) {
	ms := wrap_cloudinary2.GetMediaServer()
	result := make([]*uploader.UploadResult, 0)
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		resp, err := ms.Upload.Upload(ctx, f, uploader.UploadParams{
			Folder:      wrap_cloudinary2.GetMediaFolder(folder),
			DisplayName: file.Filename,
			EagerAsync:  api.Bool(true),
		})
		if err != nil {
			return nil, err
		}
		result = append(result, resp)
		err = f.Close()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

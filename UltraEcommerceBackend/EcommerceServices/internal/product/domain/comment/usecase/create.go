package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_cloudinary"
	ioSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/comment/storage/io"
	"mime/multipart"
	"strings"
)

func (u *commentUseCase) CreateComment(ctx context.Context, comment ioSto.CreateComment, files []*multipart.FileHeader) (uint, error) {
	commentID, err := u.commentSto.CreateComment(ctx, comment)
	if err != nil {
		return 0, err
	}
	if len(files) == 0 {
		return commentID, nil
	}

	uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, wrap_cloudinary.Comment)
	commentMedia := make([]ioSto.CreateCommentMedia, 0)
	for _, file := range uploadedFiles {

		mediaInfo := ioSto.CreateCommentMedia{
			CommentID: commentID,
			PublicID:  file.PublicID,
			MediaPath: file.URL,
			MediaType: strings.ToUpper(file.ResourceType),
		}

		commentMedia = append(commentMedia, mediaInfo)
	}

	err = u.commentSto.CreateCommentMedia(ctx, commentMedia)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

package comment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
	"mime/multipart"
)

type UseCase interface {
	ListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) (comments []io.CommentDetail, total int64, err error)
	GetCommentDetailByID(ctx context.Context, commentID uint) (io.CommentDetail, error)
	CreateComment(ctx context.Context, comment io.CreateComment, files []*multipart.FileHeader) (uint, error)
}

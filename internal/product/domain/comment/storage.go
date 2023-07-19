package comment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
)

type Storage interface {
	ListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) ([]io.CommentDetail, error)
	ListComment(ctx context.Context, filter paging.ParamsInput) ([]io.CommentDetail, error)
	CountListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) (int64, error)
	CountListComment(ctx context.Context, filter paging.ParamsInput) (int64, error)
	GetCommentDetailByID(ctx context.Context, commentID uint) (io.CommentDetail, error)
	CreateComment(ctx context.Context, comment io.CreateComment) (commentID uint, err error)
	CreateCommentMedia(ctx context.Context, media []io.CreateCommentMedia) error

	DeleteCommentByID(ctx context.Context, commentID uint) error
	RecoveryCommentByID(ctx context.Context, commentID uint) error
}

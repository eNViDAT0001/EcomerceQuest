package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
	"gorm.io/gorm"
)

func (u *commentUseCase) ListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) (comments []ioSto.CommentDetail, total int64, err error) {
	total, err = u.commentSto.CountListCommentByProductID(ctx, productID, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	comments, err = u.commentSto.ListCommentByProductID(ctx, productID, filter)
	if err != nil {
		return nil, 0, err
	}

	return comments, total, err
}

package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
)

func (u *commentUseCase) CreateComment(ctx context.Context, comment ioSto.CreateComment, medias []ioSto.CreateCommentMedia) (uint, error) {
	commentID, err := u.commentSto.CreateComment(ctx, comment)
	if err != nil {
		return 0, err
	}
	if len(medias) == 0 {
		return commentID, nil
	}

	for i, _ := range medias {
		medias[i].CommentID = commentID
	}
	err = u.commentSto.CreateCommentMedia(ctx, medias)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

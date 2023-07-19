package usecase

import "context"

func (u *commentUseCase) DeleteCommentByID(ctx context.Context, commentID uint) error {
	return u.commentSto.DeleteCommentByID(ctx, commentID)
}

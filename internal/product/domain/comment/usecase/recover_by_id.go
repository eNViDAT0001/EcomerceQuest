package usecase

import "context"

func (u *commentUseCase) RecoverCommentByID(ctx context.Context, commentID uint) error {
	return u.commentSto.RecoveryCommentByID(ctx, commentID)
}

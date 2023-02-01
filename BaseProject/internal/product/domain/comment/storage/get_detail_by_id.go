package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"gorm.io/gorm"

	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (c commentStorage) GetCommentDetailByID(ctx context.Context, commentID uint) (io.CommentDetail, error) {
	var result io.CommentDetail
	db := wrap_gorm.GetDB()
	query := db.Table(entities.Comment{}.TableName()).
		Select("Comment.*, User.name, User.avatar, JSON_ARRAYAGG( JSON_OBJECT( 'publicID', CommentMedia.public_id, 'mediaPath', CommentMedia.media_path, 'type', CommentMedia.media_type ) ) AS media").
		Joins("LEFT JOIN CommentMedia ON CommentMedia.comment_id = Comment.id").
		Joins("JOIN User ON User.id = Comment.user_id").
		Where("Comment.id = ?", commentID).
		Where("Comment.deleted_at IS NULL").
		Group("Comment.id").
		Scan(&result)

	err := query.Error
	if err != nil {
		return result, err
	}
	if query.RowsAffected < 1 {
		return result, gorm.ErrRecordNotFound
	}

	return result, nil
}

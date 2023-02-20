package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/domain/comment/storage/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/product/entities"
)

func (c commentStorage) ListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) ([]io.CommentDetail, error) {
	result := make([]io.CommentDetail, 0)
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Comment{}.TableName()).
		Select("Comment.*, User.name, User.avatar, IF(COUNT(CommentMedia.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'publicID', CommentMedia.public_id, 'mediaPath', CommentMedia.media_path, 'type', CommentMedia.media_type))) AS media").
		Joins("LEFT JOIN CommentMedia ON CommentMedia.comment_id = Comment.id").
		Joins("JOIN User ON User.id = Comment.user_id").
		Where("Comment.product_id = ?", productID).
		Where("Comment.deleted_at IS NULL").
		Group("Comment.id")

	paging_query.SetPagingQuery(&filter, entities.Comment{}.TableName(), query)

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

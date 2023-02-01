package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
)

func (c commentStorage) CountListCommentByProductID(ctx context.Context, productID uint, filter paging.ParamsInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Comment{}.TableName()).
		Select("Comment.*, IF(COUNT(CommentMedia.id) = 0, NULL, JSON_ARRAYAGG(JSON_OBJECT( 'publicID', CommentMedia.public_id, 'mediaPath', CommentMedia.media_path, 'type', CommentMedia.media_type))) AS media").
		Joins("LEFT JOIN CommentMedia ON CommentMedia.comment_id = Comment.id").
		Where("Comment.product_id = ?", productID).
		Where("Comment.deleted_at IS NULL").
		Group("Comment.id")

	paging_query.SetCountListPagingQuery(&filter, entities.Comment{}.TableName(), query)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

package entities

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/enum"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
)

type CommentMedia struct {
	wrap_gorm.SoftDeleteModel
	CommentID uint           `gorm:"column:comment_id" json:"comment_id"`
	PublicID  string         `gorm:"column:public_id" json:"public_id"`
	MediaPath string         `gorm:"column:media_path" json:"media_path"`
	MediaType enum.MediaType `gorm:"column:media_type" json:"media_type"`
}

func (CommentMedia) TableName() string {
	return "CommentMedia"
}

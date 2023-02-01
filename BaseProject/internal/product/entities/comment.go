package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Comment struct {
	wrap_gorm.SoftDeleteModel
	ProductID    uint   `gorm:"column:product_id" json:"product_id"`
	UserID       uint   `gorm:"column:user_id" json:"user_id"`
	Name         string `gorm:"column:name" json:"name"`
	Descriptions string `gorm:"column:descriptions" json:"descriptions"`
	Rating       int    `gorm:"column:rating" json:"rating"`
}

func (Comment) WithFields() []string {
	return []string{"rating", "description"}
}
func (Comment) SearchFields() []string {
	return []string{"description"}
}
func (Comment) SortFields() []string {
	return []string{"rating", "description", "id"}
}

func (Comment) TableName() string {
	return "Comment"
}

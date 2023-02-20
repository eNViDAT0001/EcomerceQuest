package entities

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/enum"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
)

type ProductMedia struct {
	wrap_gorm.SoftDeleteModel
	ProductID uint           `gorm:"column:product_id" json:"product_id"`
	PublicID  string         `gorm:"column:public_id" json:"public_id"`
	MediaPath string         `gorm:"column:media_path" json:"media_path"`
	MediaType enum.MediaType `gorm:"column:media_type" json:"media_type"`
}

func (ProductMedia) TableName() string {
	return "ProductMedia"
}

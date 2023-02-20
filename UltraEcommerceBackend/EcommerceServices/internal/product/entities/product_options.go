package entities

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
)

type ProductOption struct {
	wrap_gorm.SoftDeleteModel
	SpecificationID *uint  `gorm:"column:specification_id" json:"specification_id"`
	ProductID       uint   `gorm:"column:product_id" json:"product_id"`
	Name            string `gorm:"column:name" json:"name"`
	Price           int    `gorm:"column:price" json:"price"`
	Quantity        int    `gorm:"column:quantity" json:"quantity"`
}

func (ProductOption) TableName() string {
	return "ProductOption"
}

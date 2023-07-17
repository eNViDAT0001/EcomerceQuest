package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type OrderItem struct {
	wrap_gorm.SoftDeleteModel
	OrderID         uint   `gorm:"column:order_id" json:"order_id"`
	ProviderID      uint   `gorm:"column:provider_id" json:"provider_id"`
	ProductID       uint   `gorm:"column:product_id" json:"product_id"`
	ProductOptionID uint   `gorm:"column:product_option_id" json:"product_option_id"`
	Name            string `gorm:"column:name" json:"name"`
	Price           int    `gorm:"column:price" json:"price"`
	Option          string `gorm:"column:option" json:"option"`
	Quantity        int    `gorm:"column:quantity" json:"quantity"`
	Discount        int    `gorm:"column:discount" json:"discount"`
	Image           string `gorm:"column:image" json:"image"`
}

func (OrderItem) WithFields() []string {
	return []string{"order_id", "category_id", "product_id", "product_option_id"}
}
func (OrderItem) SearchFields() []string {
	return []string{"name", "price", "option", "quantity", "discount"}
}
func (OrderItem) SortFields() []string {
	return []string{"order_id", "category_id", "product_id", "product_option_id", "id"}
}
func (OrderItem) CompareFields() []string {
	return []string{}
}
func (OrderItem) TableName() string {
	return "OrderItem"
}

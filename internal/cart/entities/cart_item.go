package entities

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID          uint `gorm:"column:cart_id" json:"cart_id"`
	UserID          uint `gorm:"column:user_id" json:"user_id"`
	ProductID       uint `gorm:"column:product_id" json:"product_id"`
	ProductOptionID uint `gorm:"column:product_option_id" json:"product_option_id"`
	Quantity        int  `gorm:"column:quantity" json:"quantity"`
}

func (CartItem) WithFields() []string {
	return []string{"cart_id", "user_id", "product_id", "product_option_id"}
}
func (CartItem) SearchFields() []string {
	return []string{"cart_id", "user_id", "product_id", "product_option_id", "quantity"}
}
func (CartItem) SortFields() []string {
	return []string{"cart_id", "user_id", "product_id", "id"}
}

func (CartItem) TableName() string {
	return "CartItem"
}

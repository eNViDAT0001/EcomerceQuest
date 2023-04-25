package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Product struct {
	wrap_gorm.SoftDeleteModel
	ProviderID uint   `gorm:"column:provider_id" json:"provider_id"`
	CategoryID uint   `gorm:"column:category_id" json:"category_id"`
	UserID     uint   `gorm:"column:user_id" json:"user_id"`
	Name       string `gorm:"column:name" json:"name"`
	Price      int    `gorm:"column:price" json:"price"`
	Discount   int    `gorm:"column:discount" json:"discount"`
}

func (Product) WithFields() []string {
	return []string{"name", "discount", "provider_id", "user_id", "category_id", "rating"}
}
func (Product) SearchFields() []string {
	return []string{"name", "price", "id"}
}
func (Product) SortFields() []string {
	return []string{"name", "discount", "price", "created_at", "provider_id", "user_id", "category_id", "id"}
}
func (Product) TableName() string {
	return "Product"
}

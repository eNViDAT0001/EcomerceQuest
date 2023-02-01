package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ProviderID uint `gorm:"column:provider_id" json:"provider_id"`
	UserID     uint `gorm:"column:user_id" json:"user_id"`
}

func (Cart) WithFields() []string {
	return []string{"provider_id", "user_id"}
}
func (Cart) SearchFields() []string {
	return []string{"provider_id", "user_id"}
}
func (Cart) SortFields() []string {
	return []string{"provider_id", "user_id", "id"}
}

func (Cart) TableName() string {
	return "Cart"
}

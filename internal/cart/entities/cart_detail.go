package entities

import "gorm.io/datatypes"

type CartDetail struct {
	ID         uint           `gorm:"column:id" json:"id"`
	ProviderID uint           `gorm:"column:provider_id" json:"provider_id"`
	UserID     uint           `gorm:"column:user_id" json:"user_id"`
	Name       string         `gorm:"column:name" json:"name"`
	ImagePath  string         `gorm:"column:image_path" json:"image_path"`
	Items      datatypes.JSON `gorm:"column:items" json:"items"`
}

func (CartDetail) WithFields() []string {
	return []string{"provider_id", "user_id"}
}
func (CartDetail) SearchFields() []string {
	return []string{"name", "id"}
}
func (CartDetail) SortFields() []string {
	return []string{"provider_id", "user_id", "id"}
}

func (CartDetail) TableName() string {
	return "CartDetailView"
}

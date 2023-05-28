package entities

import "gorm.io/datatypes"

type CartDetail struct {
	ID                 uint           `gorm:"column:id" json:"id"`
	ProviderID         uint           `gorm:"column:provider_id" json:"provider_id"`
	ProviderName       string         `gorm:"column:name" json:"provider_name"`
	ProviderImagePath  string         `gorm:"column:image_path" json:"provider_image_path"`
	ProviderDistrictID uint           `gorm:"column:district_id" json:"provider_district_id"`
	Items              datatypes.JSON `gorm:"column:items" json:"items"`
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

package entities

type Favorite struct {
	ID         uint `gorm:"primarykey" json:"id"`
	UserID     uint `gorm:"column:user_id" json:"user_id"`
	ProviderID uint `gorm:"column:provider_id" json:"provider_id"`
}

func (Favorite) WithFields() []string {
	return []string{}
}
func (Favorite) SearchFields() []string {
	return []string{}
}
func (Favorite) SortFields() []string {
	return []string{"user_id", "provider_id", "id"}
}
func (Favorite) TableName() string {
	return "Favorite"
}

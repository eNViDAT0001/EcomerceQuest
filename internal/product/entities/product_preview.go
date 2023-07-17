package entities

import "gorm.io/datatypes"

type ProductPreview struct {
	ID                uint           `gorm:"column:id" json:"id"`
	ProviderID        uint           `gorm:"column:provider_id" json:"provider_id"`
	CategoryID        uint           `gorm:"column:category_id" json:"category_id"`
	UserID            uint           `gorm:"column:user_id" json:"user_id"`
	Name              string         `gorm:"column:name" json:"name"`
	Price             int64          `gorm:"column:price" json:"price"`
	Media             datatypes.JSON `gorm:"column:media" json:"media"`
	Discount          int            `gorm:"column:discount" json:"discount"`
	ShortDescriptions string         `gorm:"column:short_descriptions" json:"short_descriptions"`
	Rating            float32        `gorm:"column:rating" json:"rating"`
	Height            int            `gorm:"column:height" json:"height"`
	Weight            int            `gorm:"column:weight" json:"weight"`
	Length            int            `gorm:"column:length" json:"length"`
	Width             int            `gorm:"column:width" json:"width"`
}

func (ProductPreview) WithFields() []string {
	return []string{"name", "discount", "provider_id", "user_id", "category_id", "rating", "height", "weight", "length", "width"}
}
func (ProductPreview) SearchFields() []string {
	return []string{"name", "price", "id"}
}
func (ProductPreview) SortFields() []string {
	return []string{"name", "discount", "price", "created_at", "provider_id", "user_id", "category_id", "id", "height", "weight", "length", "width"}
}
func (ProductPreview) CompareFields() []string {
	return []string{}
}
func (ProductPreview) TableName() string {
	return "ProductPreview"
}

package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type ProductDescriptions struct {
	wrap_gorm.SoftDeleteModel
	ProductID        uint   `gorm:"column:product_id" json:"product_id"`
	Name             string `gorm:"column:name" json:"name"`
	PublicID         string `gorm:"column:public_id" json:"public_id"`
	DescriptionsPath string `gorm:"column:descriptions_path" json:"descriptions_path"`
}

func (ProductDescriptions) WithFields() []string {
	return []string{"name", "discount"}
}
func (ProductDescriptions) SearchFields() []string {
	return []string{"name"}
}
func (ProductDescriptions) SortFields() []string {
	return []string{"name", "discount"}
}
func (ProductDescriptions) CompareFields() []string {
	return []string{}
}
func (ProductDescriptions) TableName() string {
	return "ProductDescriptions"
}

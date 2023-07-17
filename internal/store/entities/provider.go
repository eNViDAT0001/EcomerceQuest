package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Provider struct {
	wrap_gorm.SoftDeleteModel
	UserID     uint   `gorm:"column:user_id" json:"user_id"`
	Name       string `gorm:"column:name" json:"name"`
	ImagePath  string `gorm:"column:image_path" json:"image_path"`
	Province   string `gorm:"province" json:"province"`
	District   string `gorm:"district" json:"district"`
	Ward       string `gorm:"ward" json:"ward"`
	ProvinceID int    `gorm:"province_id" json:"province_id"`
	DistrictID int    `gorm:"district_id" json:"district_id"`
	WardCode   string `gorm:"ward_code" json:"ward_code"`
	Street     string `gorm:"street" json:"street"`
}

func (Provider) WithFields() []string {
	return []string{"name", "id"}
}
func (Provider) SearchFields() []string {
	return []string{"name"}
}
func (Provider) SortFields() []string {
	return []string{"name", "id"}
}
func (Provider) CompareFields() []string {
	return []string{"created_at_>=", "created_at_<="}
}
func (Provider) TableName() string {
	return "Provider"
}

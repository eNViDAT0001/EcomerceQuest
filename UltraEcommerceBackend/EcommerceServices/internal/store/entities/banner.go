package entities

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/wrap_gorm"
	"gorm.io/datatypes"
)

type Banner struct {
	wrap_gorm.SoftDeleteModel
	UserID     uint           `gorm:"column:user_id" json:"user_id"`
	Title      string         `gorm:"column:title" json:"title"`
	Collection string         `gorm:"column:collection" json:"collection"`
	Discount   string         `gorm:"column:discount" json:"discount"`
	Image      string         `gorm:"column:image" json:"image"`
	EndTime    datatypes.Date `gorm:"column:end_time" json:"endTime"`
}

func (Banner) WithFields() []string {
	return []string{"discount", "collection", "end_time"}
}
func (Banner) SearchFields() []string {
	return []string{"discount", "collection", "end_time"}
}
func (Banner) SortFields() []string {
	return []string{"discount", "collection", "end_time", "id"}
}
func (Banner) TableName() string {
	return "Banner"
}

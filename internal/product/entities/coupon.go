package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Coupon struct {
	wrap_gorm.HardDeleteModel
	ProductID uint   `gorm:"column:product_id" json:"product_id"`
	UserID    uint   `gorm:"column:user_id" json:"user_id"`
	Name      string `gorm:"column:name" json:"name"`
	Percent   string `gorm:"column:percent" json:"percent"`
	Fixed     int    `gorm:"column:fixed" json:"fixed"`
}

func (Coupon) WithFields() []string {
	return []string{"product_id", "user_id"}
}
func (Coupon) SearchFields() []string {
	return []string{"Name", "percent", "fixed"}
}
func (Coupon) SortFields() []string {
	return []string{"percent", "Fixed", "id"}
}
func (Coupon) CompareFields() []string {
	return []string{"percent", "Fixed"}
}
func (Coupon) TableName() string {
	return "Coupon"
}

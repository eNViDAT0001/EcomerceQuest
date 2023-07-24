package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Coupon struct {
	wrap_gorm.HardDeleteModel
	UserID  uint    `gorm:"column:user_id" json:"user_id"`
	Code    string  `gorm:"column:code" json:"code"`
	Percent int     `gorm:"column:percent" json:"percent"`
	Fixed   float64 `gorm:"column:fixed" json:"fixed"`
}

func (Coupon) WithFields() []string {
	return []string{"product_id", "user_id"}
}
func (Coupon) SearchFields() []string {
	return []string{"code", "percent", "fixed"}
}
func (Coupon) SortFields() []string {
	return []string{"percent", "fixed", "id"}
}
func (Coupon) CompareFields() []string {
	return []string{"percent", "fixed"}
}
func (Coupon) TableName() string {
	return "Coupon"
}

package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type CouponDetail struct {
	wrap_gorm.HardDeleteModel
	ProductID uint `gorm:"column:product_id" json:"product_id"`
	CouponID  uint `gorm:"column:coupon_id" json:"coupon_id"`
	Total     int  `gorm:"column:total" json:"total"`
}

func (CouponDetail) WithFields() []string {
	return []string{}
}
func (CouponDetail) SearchFields() []string {
	return []string{}
}
func (CouponDetail) SortFields() []string {
	return []string{}
}
func (CouponDetail) CompareFields() []string {
	return []string{}
}
func (CouponDetail) TableName() string {
	return "CouponDetail"
}

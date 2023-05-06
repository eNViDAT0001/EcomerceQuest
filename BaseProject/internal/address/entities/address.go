package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Address struct {
	wrap_gorm.SoftDeleteModel
	UserID     *string `gorm:"column:user_id" json:"user_id"`
	Name       *string `gorm:"column:name" json:"name"`
	Gender     *bool   `gorm:"column:gender" json:"gender"`
	Phone      *string `gorm:"column:phone" json:"phone"`
	ProvinceID *int    `gorm:"column:province_id" json:"province_id"`
	DistrictID *int    `gorm:"column:district_id" json:"district_id"`
	WardCode   *string `gorm:"column:ward_code" json:"ward_code"`
	Province   *string `gorm:"column:province" json:"province"`
	District   *string `gorm:"column:district" json:"district"`
	Ward       *string `gorm:"column:ward" json:"ward"`
	Street     *string `gorm:"column:street" json:"street"`
}

func (Address) TableName() string {
	return "Address"
}

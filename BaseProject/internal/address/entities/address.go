package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Address struct {
	wrap_gorm.SoftDeleteModel
	UserID       *string `gorm:"column:user_id" json:"user_id"`
	Name         *string `gorm:"column:name" json:"name"`
	Gender       *bool   `gorm:"column:gender" json:"gender"`
	Phone        *string `gorm:"column:phone" json:"phone"`
	ProvinceCode *string `gorm:"column:province_code" json:"province_code"`
	DistrictCode *string `gorm:"column:district_code" json:"district_code"`
	WardCode     *string `gorm:"column:ward_code" json:"ward_code"`
	Street       *string `gorm:"column:street" json:"street"`
}

func (Address) TableName() string {
	return "Address"
}

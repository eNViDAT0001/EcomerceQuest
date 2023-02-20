package entities

import (
	"gorm.io/gorm"
	"time"
)

type District struct {
	Code                 *string `gorm:"primaryKey;column:code" json:"code"`
	Name                 *string `gorm:"column:name" json:"name"`
	NameEn               *string `gorm:"column:name_en" json:"name_en"`
	FullName             *string `gorm:"column:full_name" json:"full_name"`
	FullNameEn           *string `gorm:"column:full_name_en" json:"full_name_en"`
	CodeName             *string `gorm:"column:code_name" json:"code_name"`
	ProvinceCode         *string `gorm:"column:province_code" json:"province_code"`
	AdministrativeUnitID *string `gorm:"column:administrative_unit_id" json:"administrative_unit_id"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (District) TableName() string {
	return "Districts"
}

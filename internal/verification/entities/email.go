package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type Email struct {
	wrap_gorm.HardDeleteModel
	Email        string `gorm:"column:email" json:"email"`
	Descriptions string `gorm:"column:descriptions" json:"descriptions"`
}

func (Email) TableName() string {
	return "Email"
}

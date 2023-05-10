package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Payment struct {
	wrap_gorm.SoftDeleteModel
	AccountID uint   `gorm:"column:account_id" json:"account_id"`
	Email     uint   `gorm:"column:email" json:"email"`
	Name      string `gorm:"column:name" json:"name"`
	Status    *bool  `gorm:"column:status" json:"status"`
}

func (Payment) WithFields() []string {
	return []string{}
}
func (Payment) SearchFields() []string {
	return []string{}
}
func (Payment) SortFields() []string {
	return []string{}
}
func (Payment) TableName() string {
	return "Payment"
}

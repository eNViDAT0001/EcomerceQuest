package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"gorm.io/datatypes"
)

type Reservation struct {
	wrap_gorm.SoftDeleteModel
	OptionID uint           `gorm:"column:option_id" json:"option_id"`
	OrderID  uint           `gorm:"column:order_id" json:"order_id"`
	Quantity int            `gorm:"column:quantity" json:"quantity"`
	EndTime  datatypes.Time `gorm:"column:end_time" json:"end_time"`
}

func (Reservation) TableName() string {
	return "Reservation"
}

package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"gorm.io/datatypes"
)

type Order struct {
	wrap_gorm.SoftDeleteModel
	UserID             uint            `gorm:"column:user_id" json:"user_id"`
	ProviderID         uint            `gorm:"column:provider_id" json:"provider_id"`
	PaymentID          *string         `gorm:"column:payment_id" json:"payment_id"`
	PaymentURL         *string         `gorm:"column:payment_url" json:"payment_url"`
	COD                *bool           `gorm:"column:cod" json:"cod"`
	Name               string          `gorm:"column:name" json:"name"`
	Gender             *bool           `gorm:"column:gender" json:"gender"`
	Phone              string          `gorm:"column:phone" json:"phone"`
	Province           string          `gorm:"column:province" json:"province"`
	District           string          `gorm:"column:district" json:"district"`
	Ward               string          `gorm:"column:ward" json:"ward"`
	Street             string          `gorm:"column:street" json:"street"`
	Quantity           int             `gorm:"column:quantity" json:"quantity"`
	Total              int             `gorm:"column:total" json:"total"`
	Discount           int             `gorm:"column:discount" json:"discount"`
	Status             OrderStatus     `gorm:"column:status" json:"status"`
	StatusDescriptions string          `gorm:"column:status_descriptions" json:"status_descriptions"`
	DeliveredImage     *string         `gorm:"column:delivered_image" json:"delivered_image"`
	VerifyDelivered    *bool           `gorm:"column:verify_delivered" json:"verify_delivered"`
	DeliveredDate      *datatypes.Time `gorm:"column:delivered_date" json:"delivered_date"`
}

func (Order) WithFields() []string {
	return []string{"provider_id", "user_id", "name", "gender", "status", "id", "verify_delivered"}
}
func (Order) SearchFields() []string {
	return []string{"name", "gender", "status", "phone", "province", "district", "ward", "street", "quantity"}
}
func (Order) SortFields() []string {
	return []string{"delivered_date", "provider_id", "user_id", "name", "gender", "status", "total", "discount", "id", "province", "district", "ward"}
}
func (Order) TableName() string {
	return "Order"
}
func (Order) CompareFields() []string {
	return []string{"created_at_>=", "created_at_<="}
}

type OrderStatus string

const (
	WaitingOrder    OrderStatus = "WAITING"
	ConfirmedOrder  OrderStatus = "CONFIRMED"
	DeliveringOrder OrderStatus = "DELIVERING"
	DeliveredOrder  OrderStatus = "DELIVERED"
	CancelOrder     OrderStatus = "CANCEL"
)

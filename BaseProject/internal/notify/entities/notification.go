package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type Notification struct {
	wrap_gorm.SoftDeleteModel
	UserID  uint             `gorm:"column:user_id" json:"user_id"`
	Content string           `gorm:"column:content" json:"content"`
	Seen    bool             `gorm:"column:seen" json:"seen"`
	Type    NotificationType `gorm:"column:type" json:"type"`
}

func (Notification) WithFields() []string {
	return []string{"created_at", "description"}
}
func (Notification) SearchFields() []string {
	return []string{"UserID", "To"}
}
func (Notification) SortFields() []string {
	return []string{"created_at", "id"}
}

func (Notification) TableName() string {
	return "Notification"
}

type NotificationType string

const (
	Order   NotificationType = "ORDER"
	Product NotificationType = "PRODUCT"
	Cart    NotificationType = "CART"
)

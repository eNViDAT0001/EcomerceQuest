package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type Notification struct {
	wrap_gorm.SoftDeleteModel
	UserID  uint   `gorm:"column:user_id" json:"user_id"`
	Content string `gorm:"column:content" json:"content"`
	Seen    *bool  `gorm:"column:seen" json:"seen"`
	URl     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
}

func (Notification) WithFields() []string {
	return []string{"created_at", "description", "seen"}
}
func (Notification) SearchFields() []string {
	return []string{"UserID", "To", "Content"}
}
func (Notification) SortFields() []string {
	return []string{"created_at", "id"}
}

func (Notification) TableName() string {
	return "Notification"
}

type NotificationType string

package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type Message struct {
	wrap_gorm.SoftDeleteModel
	UserID   uint        `gorm:"column:user_id" json:"user_id"`
	Content  string      `gorm:"column:content" json:"content"`
	ToUserID string      `gorm:"column:to_user_id" json:"to_user_id"`
	Seen     *bool       `gorm:"column:seen" json:"seen"`
	Type     MessageType `gorm:"column:type" json:"type"`
}

func (Message) WithFields() []string {
	return []string{"created_at", "description", "seen"}
}
func (Message) SearchFields() []string {
	return []string{"UserID", "To"}
}
func (Message) SortFields() []string {
	return []string{"created_at", "id"}
}

func (Message) TableName() string {
	return "Message"
}

type MessageType string

const (
	Text  MessageType = "TEXT"
	Media MessageType = "MEDIA"
)

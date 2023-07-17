package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type Message struct {
	wrap_gorm.SoftDeleteModel
	FromUserID uint        `gorm:"column:from_user_id" json:"from_user_id"`
	Content    string      `gorm:"column:content" json:"content"`
	ToUserID   uint        `gorm:"column:to_user_id" json:"to_user_id"`
	Seen       *bool       `gorm:"column:seen" json:"seen"`
	Type       MessageType `gorm:"column:type" json:"type"`
}

func (Message) WithFields() []string {
	return []string{"seen", "user_id", "to_user_id"}
}
func (Message) SearchFields() []string {
	return []string{"description", "name"}
}
func (Message) SortFields() []string {
	return []string{"id", "created_at"}
}
func (Message) CompareFields() []string {
	return []string{}
}
func (Message) TableName() string {
	return "Message"
}

type MessageType string

const (
	Text  MessageType = "TEXT"
	Media MessageType = "MEDIA"
)

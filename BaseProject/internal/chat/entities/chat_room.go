package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type ChatRoom struct {
	wrap_gorm.SoftDeleteModel
}

func (ChatRoom) WithFields() []string {
	return []string{"seen", "user_id", "to_user_id"}
}
func (ChatRoom) SearchFields() []string {
	return []string{"description"}
}
func (ChatRoom) SortFields() []string {
	return []string{"id", "created_at"}
}

func (ChatRoom) TableName() string {
	return "ChatRoom"
}

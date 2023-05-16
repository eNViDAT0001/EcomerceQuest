package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type ChatRoom struct {
	wrap_gorm.SoftDeleteModel
}

func (ChatRoom) WithFields() []string {
	return []string{}
}
func (ChatRoom) SearchFields() []string {
	return []string{}
}
func (ChatRoom) SortFields() []string {
	return []string{}
}

func (ChatRoom) TableName() string {
	return "ChatRoom"
}

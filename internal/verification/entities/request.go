package entities

import "github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"

type RequestContact struct {
	wrap_gorm.HardDeleteModel
	Subject       string `gorm:"column:subject" json:"subject"`
	Content       string `gorm:"column:content" json:"content"`
	AttachedFiles string `gorm:"column:attached_file" json:"attached_file"`
	Type          string `gorm:"column:type" json:"type"`
	Seen          bool   `gorm:"column:seen" json:"seen"`
}

func (c RequestContact) WithFields() []string {
	return []string{"id"}
}

func (c RequestContact) SearchFields() []string {
	return []string{"id", "subject", "content", "attached_file", "type", "seen"}
}

func (c RequestContact) SortFields() []string {
	return []string{"id", "seen"}
}

func (RequestContact) TableName() string {
	return "RequestContact"
}

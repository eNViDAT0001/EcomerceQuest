package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
)

type Provider struct {
	wrap_gorm.SoftDeleteModel
	UserID    uint   `gorm:"column:user_id" json:"user_id"`
	Name      string `gorm:"column:name" json:"name"`
	ImagePath string `gorm:"column:image_path" json:"image_path"`
}

func (Provider) WithFields() []string {
	return []string{"name", "id"}
}
func (Provider) SearchFields() []string {
	return []string{"name"}
}
func (Provider) SortFields() []string {
	return []string{"name", "id"}
}
func (Provider) TableName() string {
	return "Provider"
}

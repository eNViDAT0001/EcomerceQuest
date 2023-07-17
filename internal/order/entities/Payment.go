package entities

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID        string         `gorm:"primarykey" json:"id"`
	AccountID string         `gorm:"column:account_id" json:"account_id"`
	Email     string         `gorm:"column:email" json:"email"`
	Name      string         `gorm:"column:name" json:"name"`
	Status    *bool          `gorm:"column:status" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Payment) WithFields() []string {
	return []string{}
}
func (Payment) SearchFields() []string {
	return []string{}
}
func (Payment) SortFields() []string {
	return []string{}
}
func (Payment) CompareFields() []string {
	return []string{}
}
func (Payment) TableName() string {
	return "Payment"
}

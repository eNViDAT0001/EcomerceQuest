package entities

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	wrap_gorm.SoftDeleteModel
	Username *string   `gorm:"column:username" json:"username"`
	Password *string   `gorm:"column:password" json:"password"`
	Salt     *string   `gorm:"column:salt" json:"salt"`
	Name     *string   `gorm:"column:name" json:"name"`
	Birthday *string   `gorm:"column:birthday" json:"birthday"`
	Gender   *bool     `gorm:"column:gender" json:"gender"`
	Email    *string   `gorm:"column:email" json:"email"`
	Phone    *string   `gorm:"column:phone" json:"phone"`
	Type     *UserType `gorm:"column:type" json:"type"`
	Avatar   *string   `gorm:"column:avatar" json:"avatar"`
}

func (User) WithFields() []string {
	return []string{"username", "name", "gender", "email", "phone", "type"}
}
func (User) SearchFields() []string {
	return []string{"birthday", "username", "name", "gender", "email", "phone", "type"}
}
func (User) SortFields() []string {
	return []string{"username", "name", "gender", "email", "phone", "type", "id"}
}
func (User) TableName() string {
	return "User"
}

func (u User) HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u User) ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}

type UserType string

const (
	Seller UserType = "SELLER"
	Buyer  UserType = "BUYER"
	Admin  UserType = "ADMIN"
	Banned UserType = "BANNED"
)

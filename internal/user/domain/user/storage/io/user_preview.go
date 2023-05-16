package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"

type UserPreview struct {
	ID       uint
	Username string
	Email    string
	Phone    string
	Type     entities.UserType
}

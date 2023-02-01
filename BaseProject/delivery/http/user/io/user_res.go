package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"

type UserRes struct {
	ID       uint32 
	Username string  `json:"username"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   *bool   `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Type     entities.UserType `json:"type"`
}

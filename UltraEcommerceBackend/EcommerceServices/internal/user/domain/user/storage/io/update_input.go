package io

import "github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"

type UpdateUserInput struct {
	Username string
	Name     string
	Birthday string
	Gender   *bool
	Email    string
	Phone    string
	Avatar   string
	Type     entities.UserType
}

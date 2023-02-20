package io

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
)

type CreateUserInput struct {
	ID       uint32
	Username string
	Password string
	Salt     string
	Name     string
	Birthday string
	Gender   bool
	Email    string
	Phone    string
	Type     entities.UserType
}

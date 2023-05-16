package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"

type CreateUserReq struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Name     *string `json:"name"     binding:"required"`
	Birthday *string `json:"birthday" binding:"required"`
	Gender   *bool   `json:"gender"   binding:"required"`
	Email    *string `json:"email"    binding:"required"`
	Phone    *string `json:"phone"    binding:"required"`
	Avatar   *string `json:"avatar"`
	Type     entities.UserType
}

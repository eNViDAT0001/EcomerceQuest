package io

import "github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"

type UpdateUserFullDetailReq struct {
	Username string            `json:"username"`
	Name     string            `json:"name"`
	Birthday string            `json:"birthday"`
	Gender   *bool             `json:"gender"`
	Email    string            `json:"email"`
	Phone    string            `json:"phone"`
	Avatar   string            `json:"avatar"`
	Type     entities.UserType `json:"type"`
}
type UpdateUserIdentityReq struct {
	Password string            `json:"password" binding:"required"`
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Phone    string            `json:"phone"`
	Type     entities.UserType `json:"type"`
}
type UpdateUserInfoReq struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   *bool  `json:"gender"`
}

type UpdateUserReq interface {
	UpdateUserFullDetailReq | UpdateUserIdentityReq | UpdateUserInfoReq
}

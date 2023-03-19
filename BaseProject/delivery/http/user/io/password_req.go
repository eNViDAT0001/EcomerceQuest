package io

type SetNewPasswordReq struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
type ResetPasswordReq struct {
	Email string `json:"email" binding:"required"`
}

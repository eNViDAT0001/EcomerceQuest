package io

type SetNewPasswordReq struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
type ResetPasswordReq struct {
	Code        string `json:"code"`
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

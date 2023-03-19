package io

type CodeVerificationForm struct {
	Token  string `json:"token"`
	UserID uint   `json:"user_id"`
	Code   string `json:"code"`
}

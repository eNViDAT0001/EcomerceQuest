package io

type EmailReq struct {
	Email        string `json:"email" binding:"required"`
	Descriptions string `json:"descriptions" binding:"required"`
}

package io

type EmailReq struct {
	Name  string `json:"name" `
	Email string `json:"email" binding:"required"`

	Subject      string `json:"subject"`
	Descriptions string `json:"descriptions" binding:"required"`
}
type SendEmailReq struct {
	Name  string `json:"name" `
	Email string `json:"email" binding:"required"`

	Subject      string `json:"subject"`
	Descriptions string `json:"descriptions" binding:"required"`
}

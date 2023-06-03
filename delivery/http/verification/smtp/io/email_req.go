package io

type EmailReq struct {
	Name  string `json:"name" `
	Email string `json:"email" binding:"required"`

	Subject      string `json:"subject"`
	Descriptions string `json:"descriptions" binding:"required"`

	AttachedFiles string `json:"attached_files"`
	Type          string `json:"type"`
}
type SendEmailReq struct {
	Name  string `json:"name" `
	Email string `json:"email" binding:"required"`

	Subject      string `json:"subject"`
	Descriptions string `json:"descriptions" binding:"required"`
}

package io

type EmailForm struct {
	ID          uint     `json:"id"`
	Subject     string   `json:"subject"`
	Content     string   `json:"content"`
	To          []string `json:"to"`
	Cc          []string `json:"cc"`
	Bcc         []string `json:"bcc"`
	AttachFiles []string `json:"attach_files"`
}
type CreateSmtp struct {
	ID          uint   `json:"id"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
	ToEmails    string `json:"to"`
	Cc          string `json:"cc"`
	Bcc         string `json:"bcc"`
	AttachFiles string `json:"attach_files"`
}
